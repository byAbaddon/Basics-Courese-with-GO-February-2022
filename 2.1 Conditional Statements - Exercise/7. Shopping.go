package main

import "fmt"
import "math"

func main() {
	var budget, videoCardsCount, processorsCount, ramCount float64
	fmt.Scan(&budget, &videoCardsCount, &processorsCount, &ramCount)

	video := videoCardsCount * 250
	cpu := video * 0.35 * processorsCount
	ram := video * 0.1 * ramCount
	sum := video + cpu + ram

	if videoCardsCount > processorsCount {
		sum *= 0.85
	}

	total := math.Abs(sum - budget)

	if budget >= sum {
		fmt.Printf("You have %.2f leva left!", total)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva more!", total)
	}

}

/*
name   :shopping
input  :900 2 1 3
output :You have 198.75 leva left!
*/
