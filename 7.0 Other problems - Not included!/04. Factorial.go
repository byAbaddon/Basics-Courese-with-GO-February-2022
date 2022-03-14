package main

import  "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	var f int
	fmt.Scan(&f)
	fmt.Println(factorial(f))
}

/*
name   :Factorial
input  :5
output :120
*/
