load("//tools:defs.bzl", "go_library", "go_test")

package(
    default_applicable_licenses = ["//:license"],
    licenses = ["notice"],
)

go_test(
    name = "image_test",
    size = "large",
    srcs = [
        "image_test.go",
    ],
    data = [
        "latin10k.txt",
        "mysql.sql",
        "ruby.rb",
        "ruby.sh",
    ],
    library = ":image",
    tags = [
        # Requires docker and runsc to be configured before the test runs.
        "manual",
        "local",
    ],
    visibility = ["//:sandbox"],
    deps = [
        "//pkg/test/dockerutil",
        "//pkg/test/testutil",
    ],
)

go_library(
    name = "image",
    srcs = ["image.go"],
)
