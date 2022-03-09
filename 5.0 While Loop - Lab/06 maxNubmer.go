package main

import (
	"fmt"
	"strconv"
)

func main() {
	var biggestNum = -1000000000

	for {
		var token string
		fmt.Scan(&token)

		if token == "Stop" {
			fmt.Println(biggestNum)
			break
		}

		num, _ := strconv.Atoi(token)

		if biggestNum < num {
			biggestNum = num
		}
	}

}

/*
name   :maxNumber
input  :100 99 80 70 Stop
output :100
*/
