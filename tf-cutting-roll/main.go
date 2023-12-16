package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	var guests int
	_, err := fmt.Scan(&guests) // standard input
	if err != nil {
		log.Println(err)
	}

	fmt.Println(calcCutting(guests)) // standard output
}

var ddd time.Time

func calcCutting(guests int) int {
	var result int
	if guests == 1 {
		return 0
	}

	if powerOf2(guests) {
		for i := 1; i <= guests; i++ {
			if math.Pow(2, float64(i)) == float64(guests) {
				return i
			}
		}
	} else if guests%2 != 0 {
		return (guests + 1) / 2
	} else {
		// вот тут не придумал
	}

	return result
}

func powerOf2(n int) bool {
	if n == 2 {
		return true
	}

	return n%(1<<2) == 0
}
