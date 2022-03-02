package main

import "fmt"

func main() {
	var chicken, fish, vegetable float64
	fmt.Scan(&chicken, &fish, &vegetable)
	chicken *= 10.35
	fish *= 12.40
	vegetable *= 8.15
	subtotal := chicken + fish + vegetable
	dessert := 1.20
	delivery := 2.5
	total := subtotal*dessert + float64(delivery)
	fmt.Printf("%.2f", total)
}

/*
name   :07. Food Delivery
input  :2 4 3
output :116.2
*/
