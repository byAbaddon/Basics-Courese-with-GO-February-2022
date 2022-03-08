package main

import "fmt"

func main() {
	var n int
	var max = -9999999.0
	var min = 9999999.0
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var currentNum float64
		fmt.Scan(&currentNum)

		if max < currentNum {
			max = currentNum
		}

		if min > currentNum {
			min = currentNum
		}
	}

	fmt.Println("Max number:", max)
	fmt.Println("Min number:", min)

}

/*
name   :NumberOfSequence
input  :5 10 20 304 0 50
output :
Max number: 304
Min number: 0

*/
