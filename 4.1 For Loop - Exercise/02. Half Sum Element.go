package main

import (
	"fmt"
	"math"
)

func main() {
	var n, sum, bigestNum int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var currentNum int
		fmt.Scan(&currentNum)
		if bigestNum < currentNum {
			bigestNum = currentNum
		}
		sum += currentNum
	}

	result := sum - bigestNum

	if result == bigestNum {
		fmt.Printf("Yes\nSum = %v", bigestNum)
	} else {
		fmt.Printf("No\nDiff = %v", int(math.Abs(float64(bigestNum-result))))
	}

}

/*
name   :2. Half Sum Element
input  :7 3 4 1 1 2 12 1
output :
Yes
Sum = 12
*/
