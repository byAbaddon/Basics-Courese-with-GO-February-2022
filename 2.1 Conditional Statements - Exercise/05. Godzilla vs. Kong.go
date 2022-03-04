package main

import (
	"fmt"
	"math"
)

func main() {
	var budget, statists, clothes float64
	fmt.Scan(&budget, &statists, &clothes)
	decors := budget / 10
	clothesSum := statists * clothes

	if statists > 150 {
		clothesSum -= clothesSum / 10
	}

	budget -= decors + clothesSum
	if budget >= 0 {
		fmt.Printf("Action!\nWingard starts filming with %.2f leva left.", budget)
	} else {
		budget = math.Abs(budget)
		fmt.Printf("Not enough money!\nWingard needs %.2f leva more.", budget)
	}

}

/*
name   :Godzilla vs. Kong
input  :20000 120 55.5
output :Action!
Wingard starts filming with 11340.00 leva left.
*/
