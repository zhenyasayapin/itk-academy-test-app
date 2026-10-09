package main

import (
	"io"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	originalSlice := generate()
	newOriginalSlice := generate()

	if len(originalSlice) != 10 {
		t.Fatalf("Expected length of slice 10, got %d", len(originalSlice))
	}

	if slices.Equal(originalSlice, newOriginalSlice) {
		t.Fatalf("Expected different slices, got the same")
	}
}

func TestSliceExample(t *testing.T) {
	tests := []struct {
		originalSlice []int
		expectedSlice []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
	}

	var evenSlice []int

	for _, tt := range tests {
		evenSlice = sliceExample(tt.originalSlice)

		if !slices.Equal(evenSlice, tt.expectedSlice) {
			t.Fatalf("Expected slice %v, got %v", tt.expectedSlice, evenSlice)
		}
	}
}

func TestAddElement(t *testing.T) {
	tests := []struct {
		originalSlice []int
		value         int
		expectedSlice []int
	}{
		{[]int{}, 1, []int{1}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		originalSlice := slices.Clone(tt.originalSlice)
		updatedSlice := addElement(tt.originalSlice, tt.value)

		if !slices.Equal(updatedSlice, tt.expectedSlice) {
			t.Fatalf("Expected slice %v, got %v", tt.expectedSlice, updatedSlice)
		}

		if updatedSlice[len(updatedSlice)-1] != tt.value {
			t.Fatalf("Expected value %d, got %d", tt.value, updatedSlice[len(updatedSlice)-1])
		}

		if !slices.Equal(originalSlice, tt.originalSlice) {
			t.Fatalf("Original slice was changed")
		}
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		originalSlice []int
		expectedSlice []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		copySlice := copySlice(tt.originalSlice)

		if !slices.Equal(copySlice, tt.expectedSlice) {
			t.Fatalf("Expected slice %v, got %v", tt.expectedSlice, copySlice)
		}

		copySlice[0] = 99

		if tt.originalSlice[0] == copySlice[0] {
			t.Fatalf("Changing the copied slice changed the original slice")
		}
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		originalSlice []int
		index         int
		expectedSlice []int
	}{
		{[]int{1, 2, 3}, 0, []int{2, 3}},
		{[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		originalSlice := slices.Clone(tt.originalSlice)
		updatedSlice := removeElement(tt.originalSlice, tt.index)

		if !slices.Equal(updatedSlice, tt.expectedSlice) {
			t.Fatalf("Expected slice %v, got %v", tt.expectedSlice, updatedSlice)
		}

		if !slices.Equal(originalSlice, tt.originalSlice) {
			t.Fatalf("Original slice was changed")
		}
	}
}

func TestMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()
	out, _ := io.ReadAll(r)

	os.Stdout = old

	output := string(out)

	prefixes := []string{
		"Original slice      : ",
		"Only even           : ",
		"Add element         : ",
		"Copied slice        : ",
		"Remove element      : ",
		"Original slice      : ",
	}

	for _, prefix := range prefixes {
		if !strings.Contains(output, prefix) {
			t.Fatalf("Missing expected prefix: %s", prefix)
		}
	}
}
