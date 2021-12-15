package filter

import (
	"testing"
)

func TestFilterFunc(t *testing.T) {
	vi := []int{1, 2, 3, 4, 5, 6}

	want := []int{1, 2, 3}

	if got := FilterFunc(vi, func(v int) bool { return v < 4 }); len(got) != len(want) {
		t.Errorf("should get only {1, 2, 3}")
	}
}
