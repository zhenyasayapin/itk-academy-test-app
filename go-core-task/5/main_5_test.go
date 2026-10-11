package main

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		a           []int
		b           []int
		resultBool  bool
		resultSlice []int
	}{
		{[]int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, true, []int{64, 3}},
		{[]int{65, 3, 58, 678, 64}, []int{2, 43}, false, []int{}},
		{[]int{}, []int{}, false, []int{}},
	}

	for _, tt := range tests {
		exist, diff := intersection(tt.a, tt.b)
		if exist != tt.resultBool {
			t.Errorf("Want %t, got %t", tt.resultBool, exist)
		}

		slices.Sort(tt.resultSlice)
		slices.Sort(diff)

		if !slices.Equal(tt.resultSlice, diff) {
			t.Errorf("Expected equal slices, got %v and %v", tt.resultSlice, diff)
		}
	}
}
