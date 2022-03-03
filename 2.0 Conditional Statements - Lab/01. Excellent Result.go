package main

import "fmt"

func main() {
	var grade int32
	fmt.Scan(&grade)

	if grade >= 5 {
		fmt.Println("Excellent!")
	}
}

/*
name   :Excellent Result
input  :6
output :Excellent!
*/
