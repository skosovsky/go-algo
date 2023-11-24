package main

import (
	"reflect"
	"strconv"
)

func main() {
	isPalindrome(121) //nolint:gomnd
}

// result: 20ms / 6.6Mb
func isPalindrome(x int) bool {
	xFront := []rune(strconv.Itoa(x))
	var xBack []rune

	for i := len(xFront) - 1; i >= 0; i-- {
		xBack = append(xBack, xFront[i])
	}

	return reflect.DeepEqual(xFront, xBack)
}
