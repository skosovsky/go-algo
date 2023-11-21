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
