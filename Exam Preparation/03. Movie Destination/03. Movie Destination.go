package main

import "fmt"

func main() {
	var budget, dayCount float64
	var destination, season string
	fmt.Scan(&budget, &destination, &season, &dayCount)

	var country = map[string]map[string]float64{
		"Dubai":  {"Winter": 45000.0, "Summer": 40000.0},
		"Sofia":  {"Winter": 17000.0, "Summer": 12500.0},
		"London": {"Winter": 24000.0, "Summer": 20250.0},
	}

	totalSum := country[destination][season] * dayCount

	if destination == "Dubai" {
		totalSum *= 0.70
	} else if destination == "Sofia" {
		totalSum *= 1.25
	}

	if budget < totalSum {
		fmt.Printf("The director needs %.2f leva more!", totalSum-budget)
	} else {
		fmt.Printf("The budget for the movie is enough! We have %.2f leva left!", budget-totalSum)
	}
}

/*
name   :03. Movie Destination
input  :400000 Sofia Winter 20
output :The director needs 25000.00 leva more!
*/
