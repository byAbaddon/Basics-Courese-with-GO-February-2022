package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	minNum := 10000

	for i := 0; i < n; i++ {
		var currentNum int
		fmt.Scan(&currentNum)
		if minNum > currentNum {
			minNum = currentNum
		}
	}
	fmt.Println(minNum)
}

/*
name   :Min Number
input  :3 -10 20 -30
output :
*/
