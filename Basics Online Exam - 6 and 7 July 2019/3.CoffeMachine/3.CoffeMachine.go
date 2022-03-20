package main

import "fmt"

func main() {
	var coffee, ingredients string
	var count float64
	fmt.Scan(&coffee, &ingredients, &count)

	obj := map[string]map[string]float64{
		"Espresso":   {"Without": 0.90, "Normal": 1.00, "Extra": 1.20},
		"Cappuccino": {"Without": 1.00, "Normal": 1.20, "Extra": 1.60},
		"Tea":        {"Without": 0.50, "Normal": 0.60, "Extra": 0.70},
	}

	order := obj[coffee][ingredients] * count

	if coffee == "Espresso" && count > 4 {
		order -= (order * 0.25)
	}

	if ingredients == "Without" {
		order -= (order * 0.35)
	}

	if order > 15.0 {
		order -= (order * 0.20)
	}

	fmt.Printf("You bought %v cups of %v for %.2f lv.", count, coffee, order)

}

/*
name   :3.CoffeMachine
input  :Espresso Without 10
output :You bought 10 cups of Espresso for 4.39 lv.
*/
