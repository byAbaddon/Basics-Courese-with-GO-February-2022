package main

import "fmt"
import "math"

func main() {
	var budget, card, cpu, ram float64
	fmt.Scan(&budget, &card, &cpu, &ram)
	card *= 250
	cpu *= (card * 0.35)
	ram *= (card * 0.10)
	sum := card + cpu + ram

	if card > cpu {
		sum -= (sum * 0.15)
	}
	total := math.Abs(budget - sum)

	if sum <= budget {
		fmt.Printf("You have %.2f leva left!", total)
	} else {
		fmt.Printf("Not enough money! You need %.2f leva more!", total)
	}

}

/*
name   :2. Shoping
input  :900 2 1 3
output :
*/
