package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	asrt := assert.New(t)

	q := Queue[int]{}
	q.Enqueue(3)
	asrt.Equal(1, len(q), "length should be one")
	asrt.Equal(len(q), q.Len())

	asrt.Equal(3, q.Dequeue())
	asrt.Zero(q.Len())
}

func TestStack(t *testing.T) {
	asrt := assert.New(t)

	s := Stack[int]{3}
	s.Push(4)
	asrt.Equal(2, len(s))
	asrt.Equal(4, s.Pop())
	asrt.Equal(3, s.Pop())
}
