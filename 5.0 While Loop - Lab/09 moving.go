package main

import (
	"fmt"
	"strconv"
)

func main() {
	var width, length, height float64
	fmt.Scan(&width, &length, &height)

	freeSpace := width * length * height
	cubicMt := 0.0

	for {
		var token string
		fmt.Scan(&token)

		if token == "Done" {
			if freeSpace > cubicMt {
				fmt.Printf("%v Cubic meters left.", freeSpace-cubicMt)
			}
			break
		}

		box, _ := strconv.ParseFloat(token, 8)
		cubicMt += box

		if freeSpace < cubicMt {
			fmt.Printf("No more free space! You need %v Cubic meters more.", cubicMt-freeSpace)
			break
		}

	}
}

/*
name   :moving
input  :10 10 2 20 20 20 20 122
output :No more free space! You need 2 Cubic meters more.

input  :10 1 2 4 6 Done
output :10 Cubic meters left.
*/
