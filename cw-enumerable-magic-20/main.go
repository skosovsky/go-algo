package main

import (
	"fmt"
)

func main() {
	fmt.Println(eachCons([]int{1}, 1)) //nolint:gomnd
}

func eachCons(arr []int, n int) [][]int {
	result := [][]int{}

	for i := 0; n+i <= len(arr); i++ {
		result = append(result, arr[i:n+i])
	}

	return result
}
