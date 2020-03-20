workspace(name = "prysm") # important as the ethereumapi bazel patch needs this, otherwise will not build

##############################################################################
# Go
##############################################################################

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.5/rules_go-0.18.5.tar.gz"],
    sha256 = "a82a352bffae6bee4e95f68a8d80a70e87f42c4741e6a448bec11998fcc82329",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

http_archive(
    name = "io_kubernetes_build",
    sha256 = "b84fbd1173acee9d02a7d3698ad269fdf4f7aa081e9cecd40e012ad0ad8cfa2a",
    strip_prefix = "repo-infra-6537f2101fb432b679f3d103ee729dd8ac5d30a0",
    url = "https://github.com/kubernetes/repo-infra/archive/6537f2101fb432b679f3d103ee729dd8ac5d30a0.tar.gz",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "grpc_ecosystem_grpc_gateway",
    commit = "50c55a9810a974dc5a9e7dd1a5c0d295d525f283",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
)

go_repository(
    name = "com_github_bazelbuild_buildtools",
    commit = "eb1a85ca787f0f5f94ba66f41ee66fdfd4c49b70",
    importpath = "github.com/bazelbuild/buildtools",
)

go_repository(
    name = "com_github_ghodss_yaml",
    commit = "25d852aebe32c875e9c044af3eef9c7dc6bc777f",
    importpath = "github.com/ghodss/yaml",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "51d6538a90f86fe93ac480b35f37b2be17fef232",
    importpath = "gopkg.in/yaml.v2",
)

go_repository(
    name = "com_github_urfave_cli",
    commit = "e6cf83ec39f6e1158ced1927d4ed14578fda8edb",  # v1.21.0
    importpath = "github.com/urfave/cli",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet",
    commit = "6970d62e60d86fdae3c3e510e800e8a60d755a7d",
    importpath = "github.com/wealdtech/go-eth2-wallet",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet_store_filesystem",
    commit = "1eea6a48d75380047d2ebe7c8c4bd8985bcfdeca",
    importpath = "github.com/wealdtech/go-eth2-wallet-store-filesystem",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet_types",
    commit = "af67d8101be61e7c4dd8126d2b3eba20cff5dab2",
    importpath = "github.com/wealdtech/go-eth2-wallet-types",
)

go_repository(
    name = "com_github_pkg_errors",
    commit = "614d223910a179a466c1767a985424175c39b465",  # v0.9.1
    importpath = "github.com/pkg/errors",
)

go_repository(
    name = "com_github_google_uuid",
    commit = "0cd6bf5da1e1c83f8b45653022c74f71af0538a4",  # v1.1.1
    importpath = "github.com/google/uuid",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_types",
    commit = "f9c31ddf180537dd5712d5998a3d56c45864d71f",
    importpath = "github.com/wealdtech/go-eth2-types",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet_hd",
    commit = "ce0a252a01c621687e9786a64899cfbfe802ba73",
    importpath = "github.com/wealdtech/go-eth2-wallet-hd",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet_nd",
    commit = "12c8c41cdbd16797ff292e27f58e126bb89e9706",
    importpath = "github.com/wealdtech/go-eth2-wallet-nd",
)

go_repository(
    name = "com_github_shibukawa_configdir",
    commit = "e180dbdc8da04c4fa04272e875ce64949f38bd3e",
    importpath = "github.com/shibukawa/configdir",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet_store_s3",
    commit = "1c821b5161f7bb0b3efa2030eff687eea5e70e53",
    importpath = "github.com/wealdtech/go-eth2-wallet-store-s3",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet_encryptor_keystorev4",
    commit = "0c11c07b9544eb662210fadded94f40f309d8c8f",
    importpath = "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4",
)

go_repository(
    name = "com_github_wealdtech_go_ecodec",
    commit = "7473d835445a3490e61a5fcf48fe4e9755a37957",
    importpath = "github.com/wealdtech/go-ecodec",
)

go_repository(
    name = "com_github_wealdtech_go_bytesutil",
    commit = "e564d0ade555b9f97494f0f669196ddcc6bc531d",
    importpath = "github.com/wealdtech/go-bytesutil",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_wallet_encryptor_keystorev4",
    commit = "0c11c07b9544eb662210fadded94f40f309d8c8f",
    importpath = "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4",
)

go_repository(
    name = "com_github_phoreproject_bls",
    commit = "da95d4798b09e9f45a29dc53124b2a0b4c1dfc13",
    importpath = "github.com/phoreproject/bls",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    commit = "36cc7fd7051ac4707bd56c8774825df9e8de5918",
    importpath = "github.com/aws/aws-sdk-go",
)

go_repository(
    name = "com_github_wealdtech_go_indexer",
    commit = "334862c32b1e3a5c6738a2618f5c0a8ebeb8cd51",
    importpath = "github.com/wealdtech/go-indexer",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    commit = "4def268fd1a49955bfb3dda92fe3db4f924f2285",
)

go_repository(
    name = "com_github_wealdtech_go_eth2_util",
    commit = "326ebb1755651131bb8f4506ea9a23be6d9ad1dd",
    importpath = "github.com/wealdtech/go-eth2-util",
)

go_repository(
    name = "com_github_prysmaticlabs_ethereumapis",
    commit = "25f267e475788bf8e5e01cb9d73cfd0c87020822",
    importpath = "github.com/prysmaticlabs/ethereumapis",
    patch_args = ["-p1"],
        patches = [
            "//third_party:com_github_prysmaticlabs_ethereumapis-tags.patch",
        ],
)

go_repository(
    name = "com_github_prysmaticlabs_go_ssz",
    commit = "e24db4d9e9637cf88ee9e4a779e339a1686a84ee",
    importpath = "github.com/prysmaticlabs/go-ssz",
)

go_repository(
    name = "com_github_prysmaticlabs_go_bitfield",
    commit = "dbb55b15e92f897ee230360c8d9695e2f224b117",
    importpath = "github.com/prysmaticlabs/go-bitfield",
)

go_repository(
    name = "com_github_minio_sha256_simd",
    commit = "6de4475307716de15b286880ff321c9547086fdd",  # v0.1.1
    importpath = "github.com/minio/sha256-simd",
)

go_repository(
    name = "com_github_protolambda_zssz",
    commit = "632f11e5e281660402bd0ac58f76090f3503def0",
    importpath = "github.com/protolambda/zssz",
)

go_repository(
    name = "com_github_minio_highwayhash",
    importpath = "github.com/minio/highwayhash",
    commit = "02ca4b43caa3297fbb615700d8800acc7933be98"
)

go_repository(
    name = "com_github_dgraph_io_ristretto",
    commit = "99d1bbbf28e64530eb246be0568fc7709a35ebdd",  # v0.0.1
    importpath = "github.com/dgraph-io/ristretto",
)

go_repository(
    name = "com_github_cespare_xxhash",
    commit = "d7df74196a9e781ede915320c11c378c1b2f3a1f",
    importpath = "github.com/cespare/xxhash",
)