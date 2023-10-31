package main

import "fmt"

func main() {
	var numN, numK int
	fmt.Scan(&numN, &numK) // standard input

	fmt.Print(combinationWithRepeats(numN, numK)) // standard output
}

func combinationWithRepeats(numN int, numK int) int {
	return factorial(numK+numN-1) / (factorial(numN-1) * factorial(numK))
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}

//by Python
//def combination(numN, numK):
//	return factorial(numK+numN-1) // (factorial(numN-1) * factorial(numK))
//
//def factorial(num):
//	if num <= 1:
//		return 1
//	return num * factorial(num-1)
//
//numN, numK = map(int, input().split())
//print(combination(numN, numK))
