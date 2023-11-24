package main

import "fmt"

func main() {
	fmt.Println(EvenOrOdd(55)) //nolint:gomnd
}

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
