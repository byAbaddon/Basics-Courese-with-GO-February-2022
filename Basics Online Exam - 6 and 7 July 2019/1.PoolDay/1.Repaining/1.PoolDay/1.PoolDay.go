package main

import (
	"fmt"
	"math"
)

func main() {
	var people, tax, shezlong, umbrela float64
	fmt.Scan(&people, &tax, &shezlong, &umbrela)

	enterTax := people * tax
	precentShezlong := math.Ceil(people*0.75) * shezlong
	precetnUmbrela := math.Ceil(people*0.50) * umbrela

	fmt.Printf("%.2f lv.", enterTax+precentShezlong+precetnUmbrela)
}

/*
name   :1.PoolDay
input  :21 5.50 4.40 6.20
output :254.10 lv.
*/
