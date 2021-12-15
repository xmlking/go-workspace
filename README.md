# go-workspace

Experimenting with **golang-1.18** _multi-module workspaces_


## Run

###
```bash
go test -v ./lib/...
go run ./cmd/app1/...
go run ./cmd/app2/...
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