package filter

import (
	"github.com/rs/zerolog/log"
)

func FilterFunc[T any](a []T, f func(T) bool) []T {
	var n []T
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	log.Print("hello world")
	return n
}
