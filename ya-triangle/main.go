package main

import (
	"fmt"
	"log"
)

func main() {
	var side1, side2, side3 int
	_, err := fmt.Scan(&side1, &side2, &side3)
	if err != nil {
		log.Println(err)
	} // standard input

	fmt.Print(isPossibleTriangle(side1, side2, side3)) // standard output
}

func isPossibleTriangle(side1, side2, side3 int) string {
	if side1+side2 > side3 && side1+side3 > side2 && side2+side3 > side1 {
		return "YES"
	}
	return "NO"
}
