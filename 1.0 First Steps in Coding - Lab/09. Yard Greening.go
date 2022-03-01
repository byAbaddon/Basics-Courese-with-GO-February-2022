package main

import "fmt"

func main() {
	const mtPrice = 7.61
	var mt float64
	fmt.Scan(&mt)
	var subtotal = mt * mtPrice
	price := subtotal * 0.82
	discount := subtotal - price
	fmt.Printf("The final price is: %.2f lv.\nThe discount is: %.2f lv.", price, discount)
}

/*
name   :09. Yard Greening
input  :550
output :The final price is: 3432.11 lv.
			  The discount is: 753.39 lv.
*/
