package main

import (
	"fmt"
	"log"
	"unicode"
)

const (
	fourCharacters = 4
	fiveCharacters = 5
)

func main() {
	carPlatesPackage := parseCarPlates()
	for _, carPlates := range carPlatesPackage {
		showMessage(iterateCarPlates(carPlates))
	}
}

func parseDataTemperaturePackage() [][]string { // TODO: проверить сначала что работает
	var dataTemperaturePackage [][]string
	var dateTemperatureKit []string
	var err error
	var countTemperaturePackage int

	_, err = fmt.Scan(&countTemperaturePackage)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < countTemperaturePackage; i++ {
		var countTemperatureKit int

		_, err = fmt.Scan(&countTemperatureKit)
		if err != nil {
			log.Println(err)
		}

		for i := 0; i < countTemperatureKit; i++ {
			var dataTemperature string
			_, err = fmt.Scan(&dataTemperature)
			if err != nil {
				log.Println(err)
			}

			dateTemperatureKit = append(dateTemperatureKit, dataTemperature)
		}
		dataTemperaturePackage = append(dataTemperaturePackage, dateTemperatureKit)
	}

	return dataTemperaturePackage
}

func isValidCarPlate(carPlate string) bool {
	if len(carPlate) == fourCharacters {
		character1 := rune(carPlate[0])
		character2 := rune(carPlate[1])
		character3 := rune(carPlate[2])
		character4 := rune(carPlate[3])

		if unicode.IsLetter(character1) && unicode.IsDigit(character2) && unicode.IsLetter(character3) && unicode.IsLetter(character4) {
			return true
		}
	}
	if len(carPlate) == fiveCharacters {
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
	fmt.Printf("%s\n", msg) //nolint:forbidigo // it's not debug
}
