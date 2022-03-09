package main

import "fmt"

func main() {
	currentSum := 0
	var sum int
	fmt.Scan(&sum)

	for {
		var current int
		fmt.Scan(&current)
		currentSum += current

		if sum <= currentSum {
			fmt.Println(currentSum)
			break
		}
	}

}

/*
name   :sumNumbers
input  :100 10 20 30 40
output :
*/
