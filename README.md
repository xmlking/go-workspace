# go-workspace
Experimenting with **golang-1.18** _multi-module workspaces_


## golang 1.18

### workspace commands

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

### Example 

```
.
├── README.md
├── cmd
│   ├── app1
│   │   └── go.mod
│   └── app2
│       └── go.mod
├── go.work
└── lib
    └── go.mod

```

## Reference 
- [Proposal: Multi-Module Workspaces](https://go.googlesource.com/proposal/+/master/design/45713-workspace.md)