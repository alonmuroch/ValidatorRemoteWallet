workspace(name = "alonmuroch_remotewallet")

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

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "grpc_ecosystem_grpc_gateway",
    commit = "da7a886035e25b2f274f89b6f3c64bf70a9f6780",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
)

go_repository(
    name = "com_github_bazelbuild_buildtools",
    commit = "eb1a85ca787f0f5f94ba66f41ee66fdfd4c49b70",
    importpath = "github.com/bazelbuild/buildtools",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable",
    commit = "1d89a3c832915b2314551c1d2a506874d62e53f7",  # v1.22.0
    importpath = "google.golang.org/grpc",
)

