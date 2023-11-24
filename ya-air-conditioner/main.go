package main

import "fmt"

func main() {
	var tempRoom, tempCond int
	var modeCond string
	fmt.Scan(&tempRoom, &tempCond) // standard input
	fmt.Scan(&modeCond)

	fmt.Print(detectTemp(tempRoom, tempCond, modeCond)) // standard output
}

func detectTemp(tempRoom, tempCond int, modeCond string) int {
	switch modeCond {
	case "freeze":
		if tempCond >= tempRoom {
			return tempRoom
		}
		return tempCond
	case "heat":
		if tempCond <= tempRoom {
			return tempRoom
		}
		return tempCond
	case "auto":
		return tempCond
	case "fan":
		return tempRoom
	default:
		return tempRoom
	}
}
