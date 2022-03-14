package main

import (
	"fmt"
)

func main() {
	var fuel, fuelPerLoop, loop float64
	fmt.Scan(&fuel, &fuelPerLoop, &loop)
	round := 0.0

	for i := 1; i < int(loop); i++ {
		fuel -= fuelPerLoop
		fuelPerLoop += 0.1
		round++
		if fuel <= 0 {
			break
		}
	}

	if fuel >= 0 {
		fmt.Println("Congrats! You won the race!")
	} else {
		fmt.Printf("You are out of fuel in round %.0f!", round)
	}

}

/*
name   :racing
input  :70 6.5 10
output :Congrats! You won the race!

input  :50 4.3 30
output :You are out of fuel in round 11!
*/
