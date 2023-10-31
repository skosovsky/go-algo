package main

import (
	"reflect"
	"regexp"
	"strings"
)

func main() {
	isPalindrome("0P")
}

// result: 25ms / 9.3Mb
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	re := regexp.MustCompile("[a-z0-9]")
	s = strings.Join(re.FindAllString(s, -1), "")

	xFront := []rune(s)
	xBack := []rune("")

	for i := len(xFront) - 1; i >= 0; i-- {
		xBack = append(xBack, xFront[i])
	}

	return reflect.DeepEqual(xFront, xBack)
}
