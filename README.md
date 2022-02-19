# go-workspace

Experimenting with **golang-1.18** _multi-module workspaces_

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/xmlking/go-workspace)](https://github.com/xmlking/go-workspace/blob/main/go.mod)
[![Go](https://github.com/xmlking/go-workspace/actions/workflows/go.yml/badge.svg)](https://github.com/xmlking/go-workspace/actions/workflows/go.yml)


## Install

### Working with golang 1.18 (beta)

Refer [gotip](https://github.com/golang/tools/blob/master/gopls/doc/advanced.md)

```bash
go install golang.org/dl/go1.18rc1@latest
go1.18rc1 download

cd ~/bin
ln -s /Users/<username>/sdk/go1.18rc1/bin/go gotip
```

This will build the latest beta go SDK in `/Users/<username>/sdk/go1.18beta2` <br/> 
make a link from `/Users/<username>/sdk/go1.18beta2/bin/go` to `~/bin/gotip` which is in your path.

Setting VSCode with [1.18 beta](https://github.com/golang/vscode-go/blob/master/docs/advanced.md)

in this repo, Run: `gotip work sync`

## Run

```bash
# run generate first
gotip generate ./...
```

###
```bash
# root module
gotip run ./...
gotip test -v ./... 
gotip test -v -fuzz=Fuzz ./internal
# lib module
gotip test -v ./lib/...
# app modules
gotip run ./cmd/app1/...
gotip run ./cmd/app2/...
```

## Build

```bash
gotip generate ./...
gotip build -v .
# check SBOM
gotip version -m go-workspace
# run binary
./go-workspace
```
### Workspace commands

```
$ gotip help work
Usage:

        go work <command> [arguments]

The commands are:

        edit        edit go.work from tools or scripts
        init        initialize workspace file
        sync        sync workspace build list to modules
        use         add modules to workspace file

```

```bash
gotip work sync
# `go mod` examples
gotip mod download
gotip mod graph
gotip mod tidy
gotip mod verify
gotip mod why -m github.com/ssoroka/slice
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
