package main

import (
	"fmt"
	"math"
)

func main() {
	var numA, numB, numC int
	fmt.Scan(&numA, &numB, &numC) // standard input

	fmt.Println(solveEquation(numA, numB, numC)) // standard output
}

func solveEquation(numA, numB, numC int) (result interface{}) { //nolint:gofmt
	switch {
	case numA == 0 && int(math.Pow(float64(numC), 2))-numB == 0:
		result = "MANY SOLUTIONS"
	case numA == 0 && int(math.Pow(float64(numC), 2))-numB != 0 || numC < 0:
		result = "NO SOLUTION"
	default:
		resultFloat := (math.Pow(float64(numC), 2) - float64(numB)) / float64(numA)

		if resultFloat != float64(int(resultFloat)) {
			result = "NO SOLUTION"
		} else {
			result = int(resultFloat)
		}
	}
	return result
}
