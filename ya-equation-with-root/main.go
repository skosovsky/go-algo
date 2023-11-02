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
	if numC < 0 {
		return "NO SOLUTION"
	} else if numA == 0 {
		if int(math.Pow(float64(numC), 2)) == numB {
			return "MANY SOLUTION"
		} else {
			return "NO SOLUTION"
		}
	}

	resultFloat := (math.Pow(float64(numC), 2) - float64(numB)) / float64(numA)

	if resultFloat == float64(int(resultFloat)) {
		return int(resultFloat)
	} else {
		return "NO SOLUTION"
	}
}
