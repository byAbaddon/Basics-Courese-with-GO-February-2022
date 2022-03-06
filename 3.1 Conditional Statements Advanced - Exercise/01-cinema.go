package main

import "fmt"

func main() {
	var ticket string
	var row, col float64
	fmt.Scan(&ticket, &row, &col)

	projectionType := map[string]float64{"Premiere": 12.0, "Normal": 7.5, "Discount": 5.0}
	fmt.Printf("%.2f leva", projectionType[ticket]*row*col)
}

/*
name   :cinema
input  :Premiere 10 12
output :1440.00 leva
*/
