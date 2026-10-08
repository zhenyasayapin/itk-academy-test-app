package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestGetType(t *testing.T) {
	test := []struct {
		Input        any
		ExpectedType string
	}{
		{10, "int"},
		{"go", "string"},
		{true, "bool"},
		{3.14, "float64"},
		{complex64(1 + 2i), "complex64"},
	}

	for _, c := range test {
		if getType(c.Input) != c.ExpectedType {
			t.Fatalf("Expected type %s, got %s", c.ExpectedType, getType(c.Input))
		}
	}
}

func TestImplode(t *testing.T) {
	test := []struct {
		Input          []any
		ExpectedResult string
	}{
		{[]any{10, "Golang", 3.14, true, 1 + 2i}, "10Golang3.14true(1+2i)"},
	}

	for _, c := range test {
		if implode(c.Input) != c.ExpectedResult {
			t.Fatalf("Expected result %s, got %s", c.ExpectedResult, implode(c.Input))
		}
	}
}

func TestHash(t *testing.T) {
	test := []struct {
		Input        []rune
		Salt         string
		ExpectedHash string
	}{
		{[]rune("test-data"), "go-2024", "5ccfd3732ecf09564b1082887282e3bff3045a8f6bd7b8c5747889e7b88ae774"},
	}

	for _, tt := range test {
		if tt.ExpectedHash != fmt.Sprintf("%x", hash(tt.Input, tt.Salt)) {
			t.Fatalf("Expected hash %s, got %s", tt.ExpectedHash, hash(tt.Input, tt.Salt))
		}
	}
}

func TestMain(t *testing.T) {
	expected := `int
int
int
float64
string
bool
complex64
4242423.14Golangtrue(1+2i)
53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3
`

	old := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()

	os.Stdout = old

	result, _ := io.ReadAll(r)

	if string(result) != expected {
		t.Fatalf("Expected output %s, got %s", expected, string(result))
	}
}
