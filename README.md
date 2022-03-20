# go-workspace

Experimenting with **golang-1.18** _multi-module workspaces_

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/xmlking/go-workspace)](https://github.com/xmlking/go-workspace/blob/main/go.mod)
[![Go](https://github.com/xmlking/go-workspace/actions/workflows/go.yml/badge.svg)](https://github.com/xmlking/go-workspace/actions/workflows/go.yml)


## Install

### Working with golang 1.18


```shell
brew install go 
$ go version 
go version go1.18 darwin/arm64
```

## Run

```shell
# run generate first
go generate ./...
```

###
```shell
# root module
go run ./...
go test -v ./... 
go test -v -fuzz=Fuzz ./internal
# lib module
go test -v ./lib/...
# app modules
go run ./cmd/app1/...
go run ./cmd/app2/...
```

## Build

```bash
go generate ./...
go build -v .
# check SBOM
go version -m go-workspace
# run binary
./go-workspace
```
### Workspace commands

```
$ go help work
Usage:

        go work <command> [arguments]

The commands are:

        edit        edit go.work from tools or scripts
        init        initialize workspace file
        sync        sync workspace build list to modules
        use         add modules to workspace file

```

```bash
go work sync
# `go mod` examples
go mod download
go mod graph
go mod tidy
go mod verify
go work sync
go mod why -m github.com/ssoroka/slice
```

### Project structure 

```
.
├── README.md
├── cmd
│   ├── app1
│   │   ├── go.mod
│   │   └── main.go
│   └── app2
│       ├── go.mod
│       └── main.go
├── go.work
├── go.work.sum
└── lib
    ├── binaryheap
    │   ├── binaryheap.go
    │   └── binaryheap_test.go
    ├── filter
    │   ├── filter.go
    │   └── filter_test.go
    └── go.mod

```

## Reference 
- [Proposal: Multi-Module Workspaces](https://go.googlesource.com/proposal/+/master/design/45713-workspace.md)
