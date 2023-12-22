package main

import (
	"fmt"
)

func main() {
	fmt.Println(summation(213)) //nolint:gomnd
}

func summation(n int) int {
	var sum int

	for i := n; i > 0; i-- {
		sum += i
	}

	return sum
}

func summationClever(n int) int {
	return n * (n + 1) / 2
}
