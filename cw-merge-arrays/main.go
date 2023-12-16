package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mergeArrays([]int{}, []int{})) //nolint:gomnd
}

func mergeArrays(arr1 []int, arr2 []int) []int {
	arrs := append(arr1, arr2...)
	noDouble := map[int]bool{}
	result := []int{}

	for _, v := range arrs {
		noDouble[v] = true
	}

	for k := range noDouble {
		result = append(result, k)
	}

	sort.Ints(result)

	return result
}
