package main

import "fmt"

func main() {
	var strawberriesPrice, bananas, oranges, raspberries, strawberries float64
	fmt.Scan(&strawberriesPrice, &bananas, &oranges, &raspberries, &strawberries)
	raspberriesKg := strawberriesPrice / 2
	orangesKg := raspberriesKg * 0.6
	bananasKg := raspberriesKg * 0.2

	total := raspberriesKg*raspberries +
		orangesKg*oranges +
		bananasKg*bananas +
		strawberriesPrice*strawberries
	fmt.Printf("%.2f", total)
}

/*
name   :01. Fruit Market
input  :48 10 3.3 6.5 1.7
output :333.12
*/
