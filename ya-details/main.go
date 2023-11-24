package main

import (
	"fmt"
	"log"
)

func main() {
	var n, k, m int
	_, err := fmt.Scan(&n, &k, &m) // standard input
	if err != nil {
		log.Println(err)
	}

	fmt.Println(countM(n, k, m)) // standard output
}

func countM(n, k, m int) (mCount int) {
	if m > k || k > n || m == 0 {
		return 0
	}

	for {
		if n-(k/m*m) >= 0 {
			n -= (k / m * m)
			mCount += k / m
		}

		if k > n {
			break
		}
	}

	return mCount
}
