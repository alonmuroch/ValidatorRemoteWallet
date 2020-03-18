package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"github.com/wealdtech/go-bytesutil"
	e2wallet "github.com/wealdtech/go-eth2-wallet"
	filesystem "github.com/wealdtech/go-eth2-wallet-store-filesystem"
	types "github.com/wealdtech/go-eth2-wallet-types"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	port = ":50051"

	// flags
	walletConfigPath string
)

func main() {
	app := &cli.App{
		Name: "Validator Remove Wallet Server",
		Usage: "Validator Remove Wallet Server",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name:        "configPath",
				Value:       "",
				Usage:       "file path (absolute) for key manager config file",
				Destination: &walletConfigPath,
			},
		},
		Action: func(c *cli.Context) error {
			// open wallet
			var accountsMap, err = importAccounts(walletConfigPath)
			if err != nil {
				log.Fatal(err)
				return nil
			}

			// start server
			server := &Server{
				port: port,
				accountsMap: accountsMap,
			}
			server.start()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type walletOpts struct {
	Location    		string   `json:"location"`
	Accounts    		[]string `json:"accounts"`
	Passphrases 		[]string `json:"passphrases"`
}

var walletOptsHelp = "test explaining about the wallet"

func importAccounts(path string) (map[[48]byte]types.Account, error){
	// read config file
	fileopts, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		return nil, errors.Wrap(err1, "Failed to read keymanager config file")
	}
	configFile := string(fileopts)

	opts := &walletOpts{}
	err2 := json.Unmarshal([]byte(configFile), opts)
	if err2 != nil {
		return nil, nil
	}

	// get store and wallet
	store := filesystem.New(filesystem.WithLocation(opts.Location))

	// get accounts
	accountsMap := make(map[[48]byte]types.Account)
	for _, path := range opts.Accounts {
		parts := strings.Split(path, "/")
		if len(parts[0]) == 0 {
			return nil, fmt.Errorf("did not understand account specifier %q", path)
		}
		wallet, err := e2wallet.OpenWallet(parts[0], e2wallet.WithStore(store))
		if err != nil {
			return nil, err
		}

		// get accounts
		accountSpecifier := "^.*$"
		if len(parts) > 1 && len(parts[1]) > 0 {
			accountSpecifier = fmt.Sprintf("^%s$", parts[1])
		}
		re := regexp.MustCompile(accountSpecifier)
		for account := range wallet.Accounts() {
			if re.Match([]byte(account.Name())) {
				pubKey := bytesutil.ToBytes48(account.PublicKey().Marshal())
				unlocked := false
				for _, passphrase := range opts.Passphrases {
					if err := account.Unlock([]byte(passphrase)); err != nil {
						return nil, fmt.Errorf("Failed to unlock account with one of the supplied passphrases")
					} else {
						accountsMap[pubKey] = account
						unlocked = true
						break
					}
				}
				if !unlocked {
					return nil, fmt.Errorf("Failed to unlock account with any supplied passphrase; cannot validate with this key")
				}
			}
		}
	}

	return accountsMap, nil
}