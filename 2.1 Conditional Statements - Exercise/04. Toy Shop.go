package main

import "fmt"

func main() {
	var trip, puzzle, doll, bear, minion, truck float64
	fmt.Scan(&trip, &puzzle, &doll, &bear, &minion, &truck)

	toysCount := puzzle + doll + bear + minion + truck
	toysTotal := (puzzle * 2.60) + (doll * 3) + (bear * 4.1) + (minion * 8.20) + (truck * 2)

	if toysCount >= 50 {
		toysTotal = toysTotal * 0.75
	}
	toysTotal = toysTotal * 0.90

	if toysTotal >= trip {
		fmt.Printf("Yes! %.2f lv left.", toysTotal-trip)
	} else {
		fmt.Printf("Not enough money! %.2f lv needed.", trip-toysTotal)
	}
}

/*
name   :Toy Shop
input  :40.8 20 25 30 50 10
output :
*/
