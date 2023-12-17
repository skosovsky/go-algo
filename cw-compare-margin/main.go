package main

import (
	"fmt"
)

func main() {
	fmt.Println(closeCompare(4.0, 5.0, 0.0)) //nolint:gomnd
}

func closeCompare(a, b, margin float64) int {
	sub := a - b // или math.Abs(a-b), тогда упрощение в if

	if (sub <= 0 && sub >= (0-margin)) || (sub >= 0 && sub <= margin) {
		return 0
	} else if a < b {
		return -1
	}

	return 1
}

func closeCompareClever(a, b, margin float64) int {
	switch {
	case a-b > margin:
		return 1
	case b-a > margin:
		return -1
	default:
		return 0
	}
}
