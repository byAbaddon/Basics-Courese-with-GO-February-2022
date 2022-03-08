package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0

	for i := 0; i < n; i++ {
		var currentNum int
		fmt.Scan(&currentNum)
		sum += currentNum
	}

	fmt.Print(sum)

}

/*
name   :sumNumbers
input  :2
output :30
*/
