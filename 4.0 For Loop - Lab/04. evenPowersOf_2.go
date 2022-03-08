package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := 1
	for i := 0; i <= n; i+= 2 {
		fmt.Println(s)
		s *= 2 * 2
	}

}

/*
name   :evenPowersOf_2
input  :3
output :1 4
*/
