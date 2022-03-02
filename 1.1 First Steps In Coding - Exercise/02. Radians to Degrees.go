package main

import (
	"fmt"
	"math"
)

func main() {
	var rad float32
	fmt.Scan(&rad)
	fmt.Printf("%f", rad*180/math.Pi)
}

/*
name   :02. Radians to Degrees
input  :3.1416
output :180
*/
