package main

import (
	"fmt"
	"math"
)

func main() {
	var budget, fishermen float64
	var season string
	fmt.Scan(&budget, &season, &fishermen)

	seasonPrice := map[string]float64{"Spring": 3000, "Summer": 4200, "Autumn": 4200, "Winter": 2600}
	result := seasonPrice[season]

	switch {
	case fishermen <= 6:
		result -= result * 0.1
	case fishermen >= 7 && fishermen <= 11:
		result -= result * 0.15
	case fishermen > 12:
		result -= result * 0.25
	}

	if season != "Autumn" && int(fishermen) % 2 == 0 {
		result -= result * 0.05
	}

	total := math.Abs(budget - result)

	if result > budget {
		fmt.Printf("Not enough money! You need %.2f leva.", total)
	} else {
		fmt.Printf("Yes! You have %.2f leva left.", total)
	}

}

/*
name   :fishingBoat
input  :3000 Summer 11
output :Not enough money! You need 570.00 leva.
*/
