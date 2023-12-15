package main

import (
	"fmt"
	"log"
)

func main() {
	var tariffCost, tariffSize, mbCost, trafficSize int
	_, err := fmt.Scan(&tariffCost, &tariffSize, &mbCost, &trafficSize) // standard input
	if err != nil {
		log.Println(err)
	}

	fmt.Println(calcCost(tariffCost, tariffSize, mbCost, trafficSize)) // standard output
}

func calcCost(tariffCost, tariffSize, mbCost, trafficSize int) int {
	var result int

	if trafficSize > tariffSize {
		result = tariffCost + (trafficSize-tariffSize)*mbCost
	} else {
		result = tariffCost
	}

	return result
}
