package binaryheap

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

func cmp[T constraints.Ordered](a, b T) bool {
	return a > b
}

func elements() []int {
	rand.Seed(time.Now().UnixNano())
	result := make([]int, rand.Intn(19)+1)
	for i := range result {
		result[i] = rand.Intn(100)
	}
	return result
}

func validate(t *testing.T, h *BinaryHeap[int], i int) {
	t.Helper()
	l := left(i)
	validateChildren(t, h, i, l)
	validateChildren(t, h, i, l+1)
}

func validateChildren(t *testing.T, h *BinaryHeap[int], p int, c int) {
	t.Helper()
	if c < h.Len() {
		if h.cmp(h.heap[c], h.heap[p]) {
			t.Errorf("heap invariant invalidated\n[%2d]: %d\n[%2d]: %d\n", p, c, h.heap[p], h.heap[c])
			return
		}
		validate(t, h, c)
	}
}

func TestIsEmpty(t *testing.T) {
	h := New(cmp[int])
	if !h.IsEmpty() {
		t.Fatalf("IsEmpty returned false on an empty heap")
	}
	h.Push(1)
	if h.IsEmpty() {
		t.Fatalf("IsEmpty returned true on a non-empty heap")
	}
}

func TestEmptyPeek(t *testing.T) {
	h := New(cmp[int])
	if got, ok := h.Peek(); ok || got != 0 {
		t.Errorf("Peek on empty heap\n got: (%v, %t)\nwant: (0, false)", got, ok)
	}
}

func TestEmptyPop(t *testing.T) {
	h := New(cmp[int])
	if got, ok := h.Pop(); ok || got != 0 {
		t.Errorf("Pop on empty heap\n got: (%v, %t)\nwant: (0, false)", got, ok)
	}
}

func TestLen(t *testing.T) {
	xs := elements()
	h := NewWithElements(cmp[int], xs...)
	want := len(xs)
	if got := h.Len(); got != want {
		t.Errorf("Len\n got: %d\nwant: %d", got, want)
	}
}

func TestPeekAndPop(t *testing.T) {
	xs := elements()
	h := NewWithElements(cmp[int], xs...)
	sort.Slice(xs, func(i, j int) bool { return xs[i] > xs[j] })

	for _, want := range xs {
		if got, ok := h.Peek(); !ok || got != want {
			t.Errorf("Peek\n got: (%d, %t)\nwant: (%d, true)", got, ok, want)
		}
		if got, ok := h.Pop(); !ok || got != want {
			t.Fatalf("Pop\n got: (%d, %t)\nwant: (%d, true)", got, ok, want)
		}
		validate(t, h, 0)
	}
}

func TestPush(t *testing.T) {
	h := New(cmp[int])
	for _, n := range elements() {
		h.Push(n)
		validate(t, h, 0)
	}
}

func TestPushAll(t *testing.T) {
	h := New(cmp[int])
	h.PushAll(elements()...)
	validate(t, h, 0)
}
