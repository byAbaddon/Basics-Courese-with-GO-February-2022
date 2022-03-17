package main

import "fmt"
import "math"

func main() {
	var budget, nights, price, percent, subtotal float64
	fmt.Scan(&budget, &nights, &price, &percent)

	if nights > 7 {
		subtotal = (price * 0.95) * nights
	} else {
		subtotal = (price) * nights
	}

	discount := budget * percent / 100
	total := subtotal + discount
	finalSum := budget - total

	if finalSum < 0 {
		fmt.Printf("%.2f leva needed.", math.Abs(finalSum))
	} else {
		fmt.Printf("Ivanovi will be left with %.2f leva after vacation.", finalSum)
	}
}


/*
name   :02. Family Trip
input  :800.50 8 100 2
output :Ivanovi will be left with 24.49 leva after vacation.
*/