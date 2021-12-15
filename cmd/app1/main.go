package main

import (
	"github.com/rs/zerolog/log"
	"github.com/xmlking/go-workspace/lib/filter"
)

func main() {
	vi := []int{1, 2, 3, 4, 5, 6}
	vi = filter.FilterFunc(vi, func(v int) bool {
		return v < 4
	})
	log.Print(vi)
}
