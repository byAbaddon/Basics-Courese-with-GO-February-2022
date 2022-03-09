package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var k = 1
	for {
		if k <= n {
			fmt.Println(k)
			k = k*2 + 1
		} else {
			break
		}
	}
}

/*
name   :sequence2k1
input  :17
output :1 3 7 15
*/
