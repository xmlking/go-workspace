package main

import (
	"fmt"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/ssoroka/slice"
	"github.com/tidwall/btree"
)

func main() {
	list := []string{"foo", "bar", "zee"}
	log.Print(slice.Contains(list, "foo"))
	log.Print(slice.Unique([]string{"A", "B", "C", "A", "B", "C", "B", "C", "A"}))

	s := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	// sum up strings as if they were ints.
	result := slice.Reduce(s, "0", func(acc string, i int, s string) string {
		accumulator, _ := strconv.Atoi(acc)
		current, _ := strconv.Atoi(s)
		s = strconv.Itoa(accumulator + current)
		return s
	})
	log.Print(result)

	// btree test
	// create a set
	var names btree.Set[string]

	// add some names
	names.Insert("Jane")
	names.Insert("Andrea")
	names.Insert("Steve")
	names.Insert("Andy")
	names.Insert("Janet")
	names.Insert("Andy")

	// Iterate over the maps and print each user
	names.Scan(func(key string) bool {
		log.Printf("%s\n", key)
		return true
	})
	log.Printf("\n")

	// Delete a couple
	names.Delete("Steve")
	names.Delete("Andy")

	// print the map again
	names.Scan(func(key string) bool {
		log.Printf("%s\n", key)
		return true
	})
	log.Printf("\n")
	versionInfo()
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

	for _, kv := range info.Settings {
		switch kv.Key {
		case "vcs.revision":
			fmt.Printf("Revision: %s", kv.Value)
		case "vcs.time":
			LastCommit, _ := time.Parse(time.RFC3339, kv.Value)
			fmt.Printf("LastCommit: %s", LastCommit)
		case "vcs.modified":
			DirtyBuild := kv.Value == "true"
			fmt.Printf("DirtyBuild: %v", DirtyBuild)
		}
	}

	//for _, dep := range info.Deps {
	//	fmt.Println(*dep)
	//}
	//fmt.Println(info.Settings)
	printBuildInfo()

	fmt.Println(buildInfoVersion())
}

func printBuildInfo() {
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Println("Main module:")
		printModule(&info.Main)
		fmt.Println("Dependencies:")
		for _, dep := range info.Deps {
			printModule(dep)
		}
	} else {
		fmt.Println("Built without Go modules")
	}
}

func buildInfoVersion() (string, bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", false
	}
	if info.Main.Version == "(devel)" {
		return "", false
	}
	return info.Main.Version, true
}

func printModule(m *debug.Module) {
	fmt.Printf("\t%s", m.Path)
	if m.Version != "(devel)" {
		fmt.Printf("@%s", m.Version)
	}
	if m.Sum != "" {
		fmt.Printf(" (sum: %s)", m.Sum)
	}
	if m.Replace != nil {
		fmt.Printf(" (replace: %s)", m.Replace.Path)
	}
	fmt.Println()
}
