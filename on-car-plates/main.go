package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	carPlatesPackage := parseCarPlates()
	for _, carPlates := range carPlatesPackage {
		showMessage(iterateCarPlates(carPlates))
	}
}

func parseCarPlates() []string {
	var carPlatesPackage []string
	var err error
	var countCarPlates int

	_, err = fmt.Scan(&countCarPlates)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < countCarPlates; i++ {
		var carPlates string
		_, err = fmt.Scan(&carPlates)
		if err != nil {
			log.Println(err)
		}

		carPlatesPackage = append(carPlatesPackage, carPlates)
	}

	return carPlatesPackage
}

func isValidCarPlate(carPlate string) bool {
	if len(carPlate) == 4 {
		character1 := rune(carPlate[0])
		character2 := rune(carPlate[1])
		character3 := rune(carPlate[2])
		character4 := rune(carPlate[3])

		if unicode.IsLetter(character1) && unicode.IsDigit(character2) && unicode.IsLetter(character3) && unicode.IsLetter(character4) {
			return true
		}
	}
	if len(carPlate) == 5 {
		character1 := rune(carPlate[0])
		character2 := rune(carPlate[1])
		character3 := rune(carPlate[2])
		character4 := rune(carPlate[3])
		character5 := rune(carPlate[4])

		if unicode.IsLetter(character1) && unicode.IsDigit(character2) && unicode.IsDigit(character3) && unicode.IsLetter(character4) && unicode.IsLetter(character5) {
			return true
		}
	}

	return false
}

func iterateCarPlates(carPlates string) string {
	const fourCharacters = 4
	const fiveCharacters = 5

	if len(carPlates) < fourCharacters {
		return "-"
	}

	var carPlatesList string
	var currentIdx int
	var remainCharactersCarPlates = len(carPlates)

	for remainCharactersCarPlates != 0 {
		var fourCharactersCarPlate string
		var fiveCharactersCarPlate string

		if remainCharactersCarPlates >= fourCharacters {
			fourCharactersCarPlate = carPlates[currentIdx : currentIdx+fourCharacters]
		}
		if remainCharactersCarPlates >= fiveCharacters {
			fiveCharactersCarPlate = carPlates[currentIdx : currentIdx+fiveCharacters]
		}

		if isValidCarPlate(fourCharactersCarPlate) && remainCharactersCarPlates >= fourCharacters {
			carPlatesList += fourCharactersCarPlate + " "
			currentIdx += fourCharacters
			remainCharactersCarPlates -= fourCharacters
			continue
		} else if isValidCarPlate(fiveCharactersCarPlate) && remainCharactersCarPlates >= fiveCharacters {
			carPlatesList += fiveCharactersCarPlate + " "
			currentIdx += fiveCharacters
			remainCharactersCarPlates -= fiveCharacters
			continue
		}
		return "-"
	}
	return carPlatesList
}

func showMessage(msg string) {
	fmt.Printf("%s\n", msg) //nolint:forbidigo - it's not debug
}
