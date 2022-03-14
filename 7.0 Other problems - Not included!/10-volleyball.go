package main

import (
	"fmt"
	"math"
)

func main() {
	var typeYear string
	var holidayWeekend, simpleWeekend float64
	fmt.Scan(&typeYear, &holidayWeekend, &simpleWeekend)

	weekend := (48 - simpleWeekend) * 3 / 4
	holiday := (holidayWeekend * 2.0) / 3
	allGames := weekend + holiday + simpleWeekend
	if typeYear == "leap" {
		fmt.Println(math.Floor(float64(allGames) * 115 / 100))
	} else {
		fmt.Println(math.Floor(float64(allGames)))
	}

}

/*
name   :volleyball
input  :leap 5 2
output :45
*/
