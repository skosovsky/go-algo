package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num) // standard input

	fmt.Print(permutation(5)) // standard output
}

func permutation(num int) int {
	if num <= 1 {
		return 1
	}
	return num * permutation(num-1)
}

//by Python
//def factorial(value):
//	if value <= 1:
//		return 1
//	return value * factorial(value-1)
//
//num = int(input())
//print(factorial(num))
