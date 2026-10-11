package main

import (
	"slices"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name       string
		a          []string
		b          []string
		difference []string
	}{
		{"both empty", []string{}, []string{}, []string{}},
		{"partial match", []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}, []string{"banana", "date", "fig"}, []string{"apple", "cherry", "43", "lead", "gno1"}},
		{"no match", []string{"apple", "banana", "cherry"}, []string{"orange", "sea", "place"}, []string{"apple", "banana", "cherry"}},
		{"full match", []string{"apple", "banana", "cherry"}, []string{"apple", "banana", "cherry"}, []string{}},
	}

	for _, tt := range tests {
		difference := findDifference(tt.a, tt.b)

		if len(difference) != len(tt.difference) {
			t.Errorf("got %d, want %d", len(difference), len(tt.difference))
		}

		slices.Sort(difference)
		slices.Sort(tt.difference)

		if !slices.Equal(difference, tt.difference) {
			t.Errorf("got %v, want %v", difference, tt.difference)
		}
	}
}
