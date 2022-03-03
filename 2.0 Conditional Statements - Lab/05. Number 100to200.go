package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch {
	case n < 100:
		fmt.Println("Less than 100")
	case n >= 100 && n <= 200:
		fmt.Println("Between 100 and 200")
	default:
		fmt.Println("Greater than 200")
	}
}

/*
name   :Number 100to200
input  :95
output :Less than 100
*/
