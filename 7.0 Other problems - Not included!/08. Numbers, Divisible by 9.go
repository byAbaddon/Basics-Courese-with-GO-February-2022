package main

import (
	"fmt"
	"strings"
)

func main() {
	var startN, endN int
	fmt.Scan(&startN, &endN)

	var outputNumbers []int
	sum := 0

	for i := startN; i <= endN; i++ {
		if i%9 == 0 {
			sum += i
			outputNumbers = append(outputNumbers, i)
		}
	}
	fmt.Println("The sum:", sum)
	fmt.Println(strings.Trim(fmt.Sprint(outputNumbers), "[]"))
}

/*
name   :Numbers, Divisible by 9
input  :100 200
output :
*/
