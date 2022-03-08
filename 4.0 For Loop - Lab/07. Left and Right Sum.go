package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	var leftSum, rightSum float64
	var arrNumbers []float64

	for i := 1; i <= n*2; i++ {
		var currentNum float64
		fmt.Scan(&currentNum)
		arrNumbers = append(arrNumbers, float64(currentNum))
	}

	for i := 0; i < len(arrNumbers)/2; i++ {
		leftSum += arrNumbers[i]
	}

	for i := len(arrNumbers) / 2; i < len(arrNumbers); i++ {
		rightSum += arrNumbers[i]
	}

	if leftSum == rightSum {
		fmt.Printf("Yes, sum = %v", leftSum)
	} else {
		fmt.Printf("No, diff = %d", int(math.Abs(leftSum-rightSum)))
	}

}

/*
name   :Left and Right Sum
input  :2 10 90 60 40
output :Yes, sum = 100
*/
