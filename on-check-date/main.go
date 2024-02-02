package main

import (
	"fmt"
	"log"
	"time"
)

const (
	formatDate = "02.01.2006"
)

func main() {
	datesPackage := parseDates()
	for _, kitDates := range datesPackage {
		if isValidKitDates(kitDates) {
			showMessage("YES")
		} else {
			showMessage("NO")
		}
	}
}

func parseDates() [][]int {
	var datesPackage [][]int
	var err error
	var countKitDates int

	_, err = fmt.Scan(&countKitDates)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < countKitDates; i++ {
		var day, month, year int
		var kitDates []int
		_, err = fmt.Scan(&day, &month, &year)
		if err != nil {
			log.Println(err)
		}

		kitDates = append(kitDates, day, month, year)
		datesPackage = append(datesPackage, kitDates)
	}

	return datesPackage
}

func isValidKitDates(kitDates []int) bool {
	day := kitDates[0]
	month := kitDates[1]
	year := kitDates[2]

	date := fmt.Sprintf("%02d.%02d.%d", day, month, year)
	_, err := time.Parse(formatDate, date)

	return err == nil
}

func showMessage(msg string) {
	fmt.Printf("%s\n", msg) //nolint:forbidigo - it's not debug
}
