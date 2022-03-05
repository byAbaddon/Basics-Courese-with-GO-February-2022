package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(map[bool]string{false: "invalid", true: ""}[n >= 100 && n <= 200 || n == 0])
}

/*
name   :invalidNubmer
input  :75
output :invalid
*/
