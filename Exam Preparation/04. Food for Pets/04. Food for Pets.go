package main

import (
	"fmt"
	"math"
)

func main() {
	var days, food, dogFood, catFood, cookies float64
	fmt.Scan(&days, &food)

	for i := 1; i < int(days)+1; i++ {
		var dog, cat float64
		fmt.Scan(&dog, &cat)

		if i%3 == 0 {
			dogFood += dog
			catFood += cat
			cookies += (dog + cat) / 10
			continue
		}

		dogFood += dog
		catFood += cat
	}

	allEatenFood := (dogFood + catFood) / food * 100
	eatenFood := (food - (food - (dogFood + catFood))) / 10
	eatenFoodDog := dogFood / eatenFood * 10
	eatenFoodCat := catFood / eatenFood * 10

	fmt.Printf("Total eaten biscuits: %vgr.\n", math.Round(cookies))
	fmt.Printf("%.2f%% of the food has been eaten.\n", allEatenFood)
	fmt.Printf("%.2f%% eaten from the dog.\n", eatenFoodDog)
	fmt.Printf("%.2f%% eaten from the cat.\n", eatenFoodCat)

}


/*
name   :04. Food for Pets
input  :3 1000 300 20 100 30 110 40
output :
Total eaten biscuits: 15gr.
60.00% of the food has been
eaten.
85.00% eaten from the dog.
15.00% eaten from the cat.
*/