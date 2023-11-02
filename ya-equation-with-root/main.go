package main

import (
	"fmt"
	"math"
)

func main() {
	var numA, numB, NumC int
	fmt.Scan(&numA, &numB, &NumC) // standard input

	fmt.Println(solveEquation(numA, numB, NumC)) // standard output
}

func solveEquation(numA int, numB int, numC int) (result interface{}) {
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
