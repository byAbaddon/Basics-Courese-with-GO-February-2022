package main

import "fmt"

func main() {
	var n int
	if fmt.Scan(&n); n & 1 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

/*
name   :03. Even or Odd
input  :3
output :odd
*/
