# libclang

[![PkgGoDev](https://pkg.go.dev/badge/github.com/pekim/libclang)](https://pkg.go.dev/github.com/pekim/libclang)

libclang is a Go library, providing a means to parse a compilation unit with `libclang`,
and walk the unit's Cursors.

## usage

See the Go docs for the package.
There are some minimal examples in the `_example` directory.

## pre-requisites

The `clang` dev package (with header files) needs to be installed.

```sh
# rpm-based distros
[sudo] dnf install clang-devel

# apt-based distros
[sudo] apt install clang-dev
```

## developing the library

### pre-commit hook

- install `goimports` if not already installed
  - https://pkg.go.dev/golang.org/x/tools/cmd/goimports
- install `golangci-lint` if not already installed
  - https://golangci-lint.run/usage/install/#local-installation
- install the `pre-commit` application if not already installed
  - https://pre-commit.com/index.html#install
- install pre-commit hook in this repo's workspace
  - `pre-commit install`
