package main

import "fmt"

func main() {
	var budget float64
	var season string
	fmt.Scan(&budget, &season)

	var country, place string

	switch {
	case budget <= 100:
		country = "Bulgaria"
		if season == "summer" {
			place = "Camp"
			budget -= budget * 0.70
		} else {
			place = "Hotel"
			budget -= budget * 0.30
		}

	case budget > 100 && budget <= 1000:
		country = "Balkans"
		if season == "summer" {
			place = "Camp"
			budget -= budget * 0.60
		} else {
			place = "Hotel"
			budget -= budget * 0.20
		}

	case budget > 1000:
		country = "Europe"
		place = "Hotel"
		budget -= budget * 0.10

	}
	fmt.Printf("Somewhere in %s\n%s - %.2f", country, place, budget)

}

/*
name   :journey
input  :50 summer
output :Somewhere in Bulgaria
        Camp - 15.00
*/
