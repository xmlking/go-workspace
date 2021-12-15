package main

import (
	"github.com/rs/zerolog/log"
	"github.com/xmlking/go-workspace/lib/binaryheap"
)

func main() {
	// Construct a new max heap containing ints.
	h := binaryheap.New(func(a, b int) bool { return a > b })

	// Add an int to the heap.
	h.Push(1)
	log.Print(h)

	// Add multiple ints to the heap.
	h.PushAll(2, 3)
	log.Print(h)

	// Retrieve the top item.
	a, found := h.Peek()
	log.Print(a, found)

	// Retrieve the top item and remove it from the heap.
	b, found := h.Pop()
	log.Print(b, found)
}
