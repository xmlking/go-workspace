package main

import (
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/ssoroka/slice"
	"golang.org/x/exp/maps"

	"golang.org/x/exp/slices"

	"github.com/tidwall/btree"
	"github.com/xmlking/go-workspace/internal/version"
)

func main() {

	// maps
	m := map[string]bool{
		"work": false,
		"fun":  true,
	}
	log.Print("Keys", maps.Keys(m))

	list := []string{"foo", "bar", "zee"}
	log.Print(slices.Contains(list, "foo"))
	slices.SortFunc(list, func(a, b string) bool {
		return a < b
	})
	log.Print("IsSorted?", slices.IsSorted(list))
	log.Print("slices.Compact = ", slices.Compact([]string{"A", "B", "C", "C", "A", "B", "B", "B", "C", "A"}))
	log.Print("slice.Unique = ", slice.Unique([]string{"A", "B", "C", "C", "A", "B", "B", "B", "C", "A"}))

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
	log.Info().Object("build_info", version.GetBuildInfo()).Send()
	log.Info().Msgf("build_info:%s", version.GetBuildInfo().PrettyString())
	log.Info().Msg(version.GetSoftwareBOM())
}
