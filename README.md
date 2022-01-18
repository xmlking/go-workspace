# go-workspace

Experimenting with **golang-1.18** _multi-module workspaces_

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/xmlking/go-workspace)](https://github.com/xmlking/go-workspace/blob/main/go.mod)
[![Go](https://github.com/xmlking/go-workspace/actions/workflows/go.yml/badge.svg)](https://github.com/xmlking/go-workspace/actions/workflows/go.yml)

## Run

###
```bash
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

```bash
go build -v .
go build -v -ldflags "-w -s -X main.Version=qqq -X main.BuildDate=2022-01-18T18:05:56Z -X main.GitCommit=5f16c94 -X main.GitBranch=develop -X main.GitState=dirty -X main.GitSummary=v0.1.0-74-g5f16c94-dirty" .
# check SBOM
go version -m go-workspace
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