package lists

// Stack is a naive, non-concurrent Stack implementation purely meant to reduce
// boilerplate of the most common pure-go "simple stack".
type Stack[T any] []T

// Push adds all ...T elements to the stack.
func (s *Stack[T]) Push(t ...T) {
	(*s) = append((*s), t...)
}

// Peek returns the last element pushed onto the stack without altering the
// stack.
func (s Stack[T]) Peek() T {
	return Last(s)
}

// Pop removes and returns the last element pushed onto the stack.
func (s *Stack[T]) Pop() T {
	last := len(*s) - 1
	o := (*s)[last]
	*s = (*s)[:last]
	return o
}

// Queue is a naive, non-concurrent FIFO queue purely meant to reduce
// boilerplate of the most common pure-go "simple queue."
type Queue[T any] []T

// Len returns the number of items currently enqueued.
func (q Queue[T]) Len() int {
	return len(q)
}

// Enqueue adds all the items passed onto the queue.
func (q *Queue[T]) Enqueue(t ...T) {
	*q = append((*q), t...)
}

// Empty returns whether or not the Queue is currently empty.
func (q Queue[T]) Empty() bool {
	return len(q) == 0
}

// Dqueueu returns the least recently queued item.
func (q *Queue[T]) Dequeue() T {
	o := (*q)[0]
	*q = (*q)[1:]
	return o
}

// Set is a naive, non-concurrent set implementation meant to reduce
// boilerplate of the most common pure-go "set" using map of T to empty struct.
type Set[T comparable] map[T]struct{}

// Add adds an item to the set.
func (s Set[T]) Add(t T) {
	s[t] = struct{}{}
}

// Remove deletes an item from a set.
func (s Set[T]) Remove(t T) {
	delete(s, t)
}

// Contains returns whether or not the item exists in the set.
func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}

// Items returns all items currently present in the set. If none, it will be a
// null slice.
func (s Set[T]) Items() []T {
	var items []T
	for k := range s {
		items = append(items, k)
	}
	return items
}

// SetFrom returns the items given added into a set.
func SetFrom[T comparable](ts []T) Set[T] {
	s := Set[T]{}
	for _, t := range ts {
		s.Add(t)
	}
	return s
}

// SetFromFunc takes a slice and a func that returns something comparable, and
// returns a set of the unique comparable results.
func SetFromFunc[U any, T comparable](ts []U, f func(U) T) Set[T] {
	s := Set[T]{}
	for _, t := range ts {
		s.Add(f(t))
	}
	return s
}

// Last returns the last item in the slice of any slice-like type. It will panic
func Last[T ~[]U, U any](ts T) U {
	return ts[len(ts)-1]
}
