package main

import (
	"fmt"
	"os"
)

func main() {
	var city string
	var money float64
	fmt.Scan(&city, &money)

	var result float64
	arrOfTowns := [3]string{"Sofia", "Varna", "Plovdiv"}
	obj := map[string]map[string]float64{
		"Sofia":   {"s": 5, "m": 7, "l": 8, "xl": 12},
		"Varna":   {"s": 4.5, "m": 7.5, "l": 10, "xl": 13},
		"Plovdiv": {"s": 5.5, "m": 8, "l": 12, "xl": 14.5},
	}

	for _, v := range arrOfTowns {
		if v == city && money > 0 {
			switch {
			case money <= 500:
				result = (obj[city]["s"] * money) / 100
			case money > 500 && money <= 1000:
				result = (obj[city]["m"] * money) / 100
			case money > 1000 && money <= 10000:
				result = (obj[city]["l"] * money) / 100
			case money > 10000:
				result = (obj[city]["xl"] * money) / 100
			}

			fmt.Printf("%.2f", result)
			os.Exit(0)
		}
	}
	fmt.Println("error")
}

/*
name   :tradeCommissions
input  :Sofia 1500
output :120.00
*/

/*
Plovdiv 499.99
27.50

Varna 3874.50
387.45

Kaspichan -50
error

*/
