package main

import (
	"fmt"
	"log"
)

const (
	oneDeckShip   = 1
	twoDeckShip   = 2
	threeDeckShip = 3
	fourDeckShip  = 4

	countOneDeckShip   = 4
	countTwoDeckShip   = 3
	countThreeDeckShip = 2
	countFourDeckShip  = 1
)

func main() {
	shipsPackage := parseShips()
	for _, kitShips := range shipsPackage {
		if isValidKitShips(kitShips) {
			showMessage("YES")
		} else {
			showMessage("NO")
		}
	}
}

func parseShips() [][]int {
	var shipsPackage [][]int
	var err error
	var countKitShips int

	_, err = fmt.Scan(&countKitShips)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < countKitShips; i++ {
		var ship1, ship2, ship3, ship4, ship5, ship6, ship7, ship8, ship9, ship10 int
		var kitShips []int
		_, err = fmt.Scan(&ship1, &ship2, &ship3, &ship4, &ship5, &ship6, &ship7, &ship8, &ship9, &ship10)
		if err != nil {
			log.Println(err)
		}

		kitShips = append(kitShips, ship1, ship2, ship3, ship4, ship5, ship6, ship7, ship8, ship9, ship10)
		shipsPackage = append(shipsPackage, kitShips)
	}

	return shipsPackage
}

func isValidKitShips(kitShips []int) bool {
	var sumOneDeckShip, sumTwoDeckShip, sumThreeDeckShip, sumFourDeckShip int

	for _, ship := range kitShips {
		switch ship {
		case oneDeckShip:
			sumOneDeckShip++
		case twoDeckShip:
			sumTwoDeckShip++
		case threeDeckShip:
			sumThreeDeckShip++
		case fourDeckShip:
			sumFourDeckShip++
		default:
			return false
		}
	}

	if sumOneDeckShip != countOneDeckShip || sumTwoDeckShip != countTwoDeckShip || sumThreeDeckShip != countThreeDeckShip || sumFourDeckShip != countFourDeckShip {
		return false
	}

	return true
}

func showMessage(msg string) {
	fmt.Printf("%s\n", msg) //nolint:forbidigo - it's not debug
}
