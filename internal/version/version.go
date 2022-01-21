package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

var (
	// Version is populated by govvv in compile-time.
	Version = "untouched" // "v1.23.1"
	// BuildDate is populated by govvv.
	BuildDate string // "2021-12-16T11:41:01Z"
	// GitCommit is populated by govvv.
	GitCommit string // "86ec240af8cbd1b60bcc4c03c20da9b98005b92e"
	// GitBranch is populated by govvv.
	GitBranch string
	// GitState is populated by govvv.
	GitState string // "clean"
	// GitSummary is populated by govvv.
	GitSummary string // v1.0.0, v1.0.1-5-g585c78f-dirty, fbd157c
	// GoVersion is populated by govvv.
	GoVersion string // "go1.17.5"
	// Compiler is populated by govvv.
	Compiler string // "gc"
	// Platform is populated by govvv.
	Platform string // "darwin/amd64"
)

type VersionInfo struct {
	GitVersion   string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

// VersionMsg is the message that is shown after process started.
const versionMsg = `
version     : %s
build date  : %s
go version  : %s
go compiler : %s
platform    : %s/%s
git commit  : %s
git branch  : %s
git state   : %s
git summary : %s
`

// GetBuildInfo helper
func GetBuildInfo() string {
	return fmt.Sprintf(versionMsg, Version, BuildDate, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH,
		GitCommit, GitBranch, GitState, GitSummary)
}

func versionInfo() {
	//info, ok := debug.ReadBuildInfo()
	//if !ok {
	//	fmt.Println("Build info not found")
	//	os.Exit(1)
	//}
	//op, err := json.MarshalIndent(info.Settings, "", " ")
	//if err != nil {
	//	panic(fmt.Sprintf("error marshalling: %v", err))
	//}
	//fmt.Println(string(op))

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	fmt.Println(info)
	fmt.Println(info.GoVersion)
	fmt.Println(info.Path)
	fmt.Println(info.Main)
	fmt.Println("version: " + info.Main.Version)
	fmt.Println(info.Main.Path)
	fmt.Println(info.Main.Sum)
	fmt.Println(info.Main.Replace)
	for _, dep := range info.Deps {
		fmt.Println(*dep)
	}
	fmt.Println(info.Settings)
}

func GetSBOM() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}
	for _, dep := range info.Deps {
		fmt.Println(*dep)
	}
}
