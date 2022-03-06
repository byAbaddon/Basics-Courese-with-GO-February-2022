package main

import "fmt"

func main() {
	var degrees int64
	var partOfDay string
	fmt.Scan(&degrees, &partOfDay)

	data := map[string]map[string]string{
		"m": {"Morning": "Sweatshirt and Sneakers", "Afternoon": "Shirt and Moccasins", "Evening": "Shirt and Moccasins"},
		"a": {"Morning": "Shirt and Moccasins", "Afternoon": "T-Shirt and Sandals", "Evening": "Shirt and Moccasins"},
		"e": {"Morning": "T-Shirt and Sandals", "Afternoon": "Swim Suit and Barefoot", "Evening": "Shirt and Moccasins"},
	}

	if degrees >= 10 && degrees <= 18 {
		fmt.Printf("It's %d degrees, get your %s.", degrees, data["m"][partOfDay])
	}
	if degrees > 18 && degrees <= 24 {
		fmt.Printf("It's %d degrees, get your %s.", degrees, data["a"][partOfDay])
	}
	if degrees >= 25 {
		fmt.Printf("It's %d degrees, get your %s.", degrees, data["e"][partOfDay])
	}

}

/*
name   :summerOutfit
input  :16 Morning
output :It"s 16 degrees, get your Sweatshirt and Sneakers.
*/
