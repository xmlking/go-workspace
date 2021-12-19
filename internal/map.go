package internal

// GroupByFunc takes any slice of T and a func that returns a comparable and
// value. The map will contain all Values returned by the func.
func GroupByFunc[T, V any, K comparable](ts []T, f func(T) (K, V)) map[K][]V {
	m := map[K][]V{}
	for _, t := range ts {
		k, v := f(t)
		m[k] = append(m[k], v)
	}
	return m
}

// MapOfFunc takes any slice of T and a func that returns a comparable and
// value. The map will contain the most recent (furthest into the slice) values
// returned for a given key.
func MapOfFunc[T, V any, K comparable](ts []T, f func(T) (K, V)) map[K]V {
	m := map[K]V{}
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}
