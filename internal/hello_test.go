package internal

import (
	"testing"
)

func TestMinMax(t *testing.T) {
	max := Max(1, 5)
	min := Min(5, 1)
	if min > max {
		t.Errorf("min: %v shoud be less then max: %v", min, max)
	}
}

func FuzzMinMax(f *testing.F) {
	// Seed the initial corpus
	f.Add(78, 66)

	// Run the fuzz test
	f.Fuzz(func(t *testing.T, a int, b int) {
		min := Min(a, b)
		max := Max(a, b)
		if min > max {
			t.Errorf("min: %v shoud be less then max: %v", min, max)
		}
	})
}
