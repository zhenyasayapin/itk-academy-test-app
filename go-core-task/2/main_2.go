package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	originalSlice := generate()

	fmt.Printf("%-20s: %v\n", "Original slice", originalSlice)
	fmt.Printf("%-20s: %v\n", "Only even", sliceExample(originalSlice))
	fmt.Printf("%-20s: %v\n", "Add element", addElement(originalSlice, rand.Intn(100)))
	fmt.Printf("%-20s: %v\n", "Copied slice", copySlice(originalSlice))
	fmt.Printf("%-20s: %v\n", "Remove element", removeElement(originalSlice, 0))
	fmt.Printf("%-20s: %v\n", "Original slice", originalSlice)
}

func sliceExample(s []int) []int {
	var newSlice []int

	for _, n := range s {
		if n%2 == 0 {
			newSlice = append(newSlice, n)
		}
	}

	return newSlice
}

func addElement(s []int, n int) []int {
	return append(s, n)
}

func copySlice(s []int) []int {
	return slices.Clone(s)
}

func removeElement(s []int, i int) []int {
	if i >= len(s) || i < 0 {
		return s
	}

	clone := slices.Clone(s)

	return append(clone[:i], clone[i+1:]...)
}

func generate() []int {
	var originalSlice []int
	for range 10 {
		originalSlice = append(originalSlice, rand.Intn(100))
	}

	return originalSlice
}
