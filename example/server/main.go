package main

import (
	"log"
	"github.com/urfave/cli"
	"os"
)

var (
	port = ":50051"
	keys = [...]string{
		"959892502eb114ba7ee20b05c45181e0429256f0c6b366a84a64ea7d4c23f4ab338116fea1bd79d3821dfe5fea825db9",
		"8b1d8ccf22e269a082ab9b9d19cfd162a841a9c0642b383661fa47dee7f0d81c74ea2c8e192538c4e9b60a26250ee4c3",
	}
)

func main() {
	app := &cli.App{
		Name: "Validator Remove Wallet Server",
		Usage: "Validator Remove Wallet Server",
		Action: func(c *cli.Context) error {
			server := &Server{
				port: port,
				keys: keys,
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

