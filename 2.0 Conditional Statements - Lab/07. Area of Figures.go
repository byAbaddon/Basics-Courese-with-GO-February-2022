package main

import (
	"fmt"
	"math"
)

func main() {
	var result float64
	var figure string
	fmt.Scan(&figure)
	if figure == "square" || figure == "circle" {
		var a float64
		fmt.Scan(&a)
		if figure == "square" {
			result = math.Pow(a, 2)
		} else {
			result = a * a * math.Pi
		}
	} else {
		var a, b float64
		fmt.Scan(&a, &b)
		if figure == "rectangle" {
			result = a * b
		} else {
			result = a * b / 2
		}
	}
	fmt.Printf("%.3f", result)
}

/*
name   :Area of Figures
input  :square 5
output :25.000
*/
