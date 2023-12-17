package main

import (
	"fmt"
)

func main() {
	fmt.Println(well([]string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"})) //nolint:gomnd
}

func well(x []string) string {
	var goodCount int

	for _, v := range x {
		if goodCount > 2 {
			break
		}

		if v == "good" {
			goodCount++
		}
	}

	switch goodCount {
	case 1, 2:
		return "Publish!"
	case 3:
		return "I smell a series!"
	default:
		return "Fail!"
	}
}
