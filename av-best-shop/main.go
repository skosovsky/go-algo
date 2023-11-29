package main

import (
	"fmt"
	"time"
)

type Visit struct {
	ShopName    string
	VisitedFrom time.Time
	VisitedTo   time.Time
}

func main() {
	vasyaVisits := [][]Visit{
		{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 11, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 12, 0, 0, 0, time.UTC)}},
		{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 10, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 11, 0, 0, 0, time.UTC)}},
	}

	petyaVisits := [][]Visit{
		{{ShopName: "B", VisitedFrom: time.Date(2007, time.July, 1, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 1, 11, 1, 0, 0, time.UTC)}},
		{{ShopName: "A", VisitedFrom: time.Date(2007, time.July, 2, 8, 0, 0, 0, time.UTC), VisitedTo: time.Date(2007, time.July, 2, 10, 0, 1, 0, time.UTC)}},
	}

	fmt.Println(BestInterestingShop(vasyaVisits, petyaVisits))

}

func BestInterestingShop(vasyaVisits [][]Visit, petyaVisits [][]Visit) string {
	// Конечно я бы все это разнес по функциям и сделаем более абстрактно, но в 1 функции в этом смысла нет
	dayCount := len(vasyaVisits)
	shopCount := len(vasyaVisits[0])
	bestShop := ""
	bestShopTime := 0

	for i := 0; i < dayCount; i++ {
		// Переложим в мапу, для более удобного обращения к магазинам
		vasyaVisitsDay := make(map[string]Visit, shopCount)
		for _, val := range vasyaVisits[i] {
			vasyaVisitsDay[val.ShopName] = val
		}
		petyaVisitsDay := make(map[string]Visit, shopCount)
		for _, val := range petyaVisits[i] {
			petyaVisitsDay[val.ShopName] = val
		}

		// Сравниваем периоды по каждому магазину
		for vShop := range vasyaVisitsDay {
			for pShop := range petyaVisitsDay {
				if vShop == pShop {
					dateStart := vasyaVisitsDay[vShop].VisitedFrom.Truncate(24 * time.Hour)
					vFrom := int(vasyaVisitsDay[vShop].VisitedFrom.Sub(dateStart) / time.Nanosecond)
					vTo := int(vasyaVisitsDay[vShop].VisitedTo.Sub(dateStart) / time.Nanosecond)
					pFrom := int(petyaVisitsDay[pShop].VisitedFrom.Sub(dateStart) / time.Nanosecond)
					pTo := int(petyaVisitsDay[pShop].VisitedTo.Sub(dateStart) / time.Nanosecond)

					if vFrom == pFrom && vTo == pTo { // полное совпадения
						if bestShopTime < vTo-vFrom {
							bestShop = vShop
							bestShopTime = vTo - vFrom
						}
					} else if vFrom >= pFrom && vTo >= pTo { // совпадение слева
						if bestShopTime < pTo-vFrom {
							bestShop = vShop
							bestShopTime = pTo - vFrom
						}
					} else if vFrom <= pFrom && vTo <= pTo { // совпадение справа
						if bestShopTime < vTo-pFrom {
							bestShop = vShop
							bestShopTime = vTo - pFrom
						}
					} else if vFrom <= pFrom && vTo >= pTo { // полное вхождение справа
						if bestShopTime < pTo-pFrom {
							bestShop = vShop
							bestShopTime = pTo - pFrom
						}
					} else if vFrom >= pFrom && vTo <= pTo { // полное вхождение слева
						if bestShopTime < vTo-vFrom {
							bestShop = vShop
							bestShopTime = vTo - vFrom
						}
					} else if (vFrom >= pFrom && vTo >= pTo) || (pFrom >= vFrom && pTo >= vTo) {
						// нет вхождения
					}
				}
			}
		}
	}

	if bestShopTime == 0 {
		return "None"
	}
	return bestShop
}
