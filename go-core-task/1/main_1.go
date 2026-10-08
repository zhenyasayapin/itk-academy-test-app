package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатиричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"        // Тип string
var isActive bool = true          // Тип bool
var complexNum complex64 = 1 + 2i // Тип complex64

var vars = []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

const salt = "go-2024"

func main() {
	for _, v := range vars {
		fmt.Println(getType(v))
	}

	concat := implode(vars)
	fmt.Println(concat)

	r := []rune(concat)
	fmt.Printf("%x\n", hash(r, salt))
}

func getType(value any) string {
	return fmt.Sprintf("%T", value)
}

func implode(vars []any) string {
	var values []string

	for _, v := range vars {
		values = append(values, fmt.Sprint(v))
	}

	return strings.Join(values, "")
}

func hash(r []rune, s string) []byte {
	b := []byte(string(r))
	mid := len(b) / 2

	h := sha256.New()
	h.Write(b[:mid])
	h.Write([]byte(s))
	h.Write(b[mid:])

	return h.Sum(nil)
}
