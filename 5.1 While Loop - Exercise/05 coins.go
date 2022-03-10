package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	coins := []float64{2, 1, 0.50, 0.20, 0.10, 0.05, 0.02, 0.01}
	var counter, money float64
	fmt.Scan(&money)

	for _, coin := range coins {
		counter += math.Floor(money / coin)
		result := fmt.Sprintf("%.2f", math.Mod(money, coin))
		digit, _ := strconv.ParseFloat(result, 8)
		money = digit
	}

	fmt.Println(counter)
}

/*
name   :coins
input  :1.23
output :4
*/
