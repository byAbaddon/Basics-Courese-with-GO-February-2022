package main

import "fmt"

func main() {
	var nailon, pain, razred, hours float64
	fmt.Scan(&nailon, &pain, &razred, &hours)

	klNailon := (nailon + 2) * 1.50      
	klPain := (pain * 110 / 100) * 14.50 
	klRazred := razred * 5.00
	subtotal := klNailon + klPain + klRazred + 0.40
	total := (subtotal*0.30) * hours + subtotal

	fmt.Printf("Total expenses: %.2f lv.", total)
}

/*
name   :repainting
input  :10 11 4 8
output :Total expenses: 727.09 lv.
*/
