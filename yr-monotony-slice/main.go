package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var userSlice []int
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = text[:len(text)-1]

	userSliceString := strings.Split(text, " ")

	for _, val := range userSliceString {
		valInt, _ := strconv.Atoi(val)
		userSlice = append(userSlice, valInt)
	}

	fmt.Println(checkMonotopy(userSlice))
}

func checkMonotopy(userSlice []int) bool {
	if len(userSlice) <= 2 {
		return true
	}

	firstDirect := ""

	if userSlice[1] > userSlice[0] {
		firstDirect = "up"
	} else if userSlice[1] == userSlice[0] {
		firstDirect = "qual"
	} else {
		firstDirect = "down"
	}

	counterFalse := 0

	for i := 2; i <= len(userSlice)-1; i++ {
		if counterFalse != 0 {
			return false
		}
		seconDirect := ""

		if userSlice[i] > userSlice[i-1] {
			seconDirect = "up"
		} else if userSlice[i] == userSlice[i-1] {
			seconDirect = "qual"
		} else {
			seconDirect = "down"
		}

		if firstDirect != seconDirect {
			counterFalse++
		}
	}

	return counterFalse == 0
}

func checkSlice(slice []int) bool {
	if len(slice) <= 1 {
		return true
	}

	iterationSize := slice[0] - slice[1]
	for i := 1; i < len(slice)-1; i++ {
		if iterationSize != slice[i]-slice[i+1] {
			return false
		}
	}

	return true
}

func isMonotonic(in []int) bool {
	less := false
	great := true

	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] && !less {
			great = false
		} else if in[i] < in[i-1] && !great {
			return false
		} else if in[i] < in[i-1] && great {
			less = true
		} else if in[i] > in[i-1] && less {
			return false
		}
	}
	return true
}
