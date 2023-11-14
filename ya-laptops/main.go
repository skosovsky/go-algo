package main

import (
	"fmt"
)

func main() {
	var a1, b1, a2, b2 int
	fmt.Scan(&a1, &b1, &a2, &b2) // standard input
	//a1, b1, a2, b2 = 5, 7, 3, 2

	fmt.Println(calcSizeTable(a1, b1, a2, b2)) // standard output
}

func calcSizeTable(a1, b1, a2, b2 int) (a, b int) {
	opt1 := [2]int{b1 + b2, max(a1, a2)}
	opt2 := [2]int{a1 + b2, max(b1, a2)}
	opt3 := [2]int{max(b1, b2), a1 + a2}
	opt4 := [2]int{max(a1, b2), b1 + a2}

	minSqare := min(opt1[0]*opt1[1], opt2[0]*opt2[1], opt3[0]*opt3[1], opt4[0]*opt4[1])

	if minSqare == opt1[0]*opt1[1] {
		a = opt1[0]
		b = opt1[1]
	} else if minSqare == opt2[0]*opt2[1] {
		a = opt2[0]
		b = opt2[1]
	} else if minSqare == opt3[0]*opt3[1] {
		a = opt3[0]
		b = opt3[1]
	} else if minSqare == opt4[0]*opt4[1] {
		a = opt4[0]
		b = opt4[1]
	}

	return a, b
}
