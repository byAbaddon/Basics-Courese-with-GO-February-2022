package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	for n > 0 {
		fmt.Println(n)
		n--
	}

}

/*
name   :numbersN
input  :2
output :2 1
*/
