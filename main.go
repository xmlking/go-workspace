package main

import (
	"strconv"

	"github.com/rs/zerolog/log"
	"github.com/ssoroka/slice"
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
}
