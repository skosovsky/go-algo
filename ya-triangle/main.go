package main

import "fmt"

func main() {
	var side1, side2, side3 int
	fmt.Scan(&side1, &side2, &side3) // standard input

	fmt.Print(isPossibleTriangle(side1, side2, side3)) // standard output
}

func isPossibleTriangle(side1 int, side2 int, side3 int) string {
	if side1+side2 > side3 && side1+side3 > side2 && side2+side3 > side1 {
		return "YES"
	}
	return "NO"
}
