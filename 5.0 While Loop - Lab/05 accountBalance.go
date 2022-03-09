package main

import "fmt"

func main() {
	var sum = 0.0
	for {
		var current float64
		fmt.Scan(&current)

		if current < 0 {
			fmt.Print("Invalid operation!", "\n")
			break
		} else if current > 0 {
			sum += current
			fmt.Printf("Increase: %.2f\n", current)
		} else {
			break
		}

	}

	fmt.Printf("Total: %.2f", sum)
}

/*
name   :accountBalance
input  :5.51 69.42 100 NoMoreMoney
output :
*/
