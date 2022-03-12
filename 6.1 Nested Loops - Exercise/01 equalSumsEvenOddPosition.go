package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNum, secondNum int
	fmt.Scan(&firstNum, &secondNum)
	runes := map[int]int{48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9}

	for n := firstNum; n <= secondNum; n++ {
		odd := 0
		even := 0
		currentNum := n
		for i, v := range strconv.Itoa(n) {
			digit := runes[int(v)]

			if i%2 == 0 {
				even += digit
			} else {
				odd += digit
			}

		}

		if even == odd {
			fmt.Print(currentNum, " ")
		}

	}
}

/*
name   :equalSumsEvenOddPosition
input  :100000 100050
output :100001 100012 100023 100034 100045
*/
