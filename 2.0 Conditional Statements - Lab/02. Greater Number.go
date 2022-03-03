package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Print(math.Max(a, b))
}

/*
name   :02. Greater Number
input  :5 3
output :5
*/

