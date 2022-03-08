package main

import "fmt"
import "math"

func main() {
	var n int
	fmt.Scan(&n)

	var evenSum, oddSum float64
	var listNum []float64

	for n > 0 {
		n--
		var currentNum float64
		fmt.Scan(&currentNum)

		listNum = append(listNum, currentNum)
	}

	for i := 0; i < len(listNum); i++ {
		if i%2 == 0 {
			evenSum += listNum[i]
		} else {
			oddSum += listNum[i]
		}
	}

	if evenSum == oddSum {
		fmt.Printf("Yes\nSum = %v", evenSum)
	} else {
		fmt.Printf("No\nDiff = %v", math.Abs(evenSum-oddSum))
	}

}

/*
name   :
input  :4 10 50 60 20
output :
Yes
Sum = 70
*/
