package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(Capitalize("abcdef", []int{1, 2, 5})) //nolint:gomnd
}

func Capitalize(st string, arr []int) string {
	var line = []rune(st)
	var newline []rune
	digits := map[int]bool{}

	for _, v := range arr {
		if v > len(line) {
			continue
		}
		digits[v] = true
	}

	for i, _ := range line {
		symbol := line[i]
		if digits[i] {
			newline = append(newline, unicode.ToUpper(symbol))
		} else {
			newline = append(newline, symbol)
		}
	}

	return string(newline)
}
