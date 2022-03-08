package main

import (
	"fmt"
	"math"
)

func main() {
	var age, landryPrice, toysPrice, moneyCount, toysCount, brotherRecket float64
	fmt.Scan(&age, &landryPrice, &toysPrice)

	for i := 1; i < int(age)+1; i++ {
		if i%2 == 0 {
			moneyCount += float64(10 * i)
			brotherRecket += 1
		} else {
			toysCount++
		}
	}

	toysCount *= toysPrice
	subtotal := (toysCount + (moneyCount / 2)) - brotherRecket
	total := math.Abs(subtotal - landryPrice)

	if subtotal >= landryPrice {
		fmt.Printf("Yes! %.2f", total)
	} else {
		fmt.Printf("No! %.2f", total)
	}
}

/*
name   :Clever Lily
input  :10 170.00 6
output :Yes! 5.00
*/
