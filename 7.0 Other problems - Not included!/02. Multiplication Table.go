package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Println(i, "*", n, "=", i*n)
	}

}

/*
name   :Multiplication Table
input  :1
output :1 * 1 = 1.....
*/
