package main

import (
	"fmt"
	"strconv"
)

func main() {
	var minNum = 1000000000

	for {
		var token string
		fmt.Scan(&token)

		if token == "Stop" {
			fmt.Println(minNum)
			break
		}

		num, _ := strconv.Atoi(token)

		if minNum > num {
			minNum = num
		}
	}

}

/*
name   :minNumber
input  :100 99 80 70 Stop
output :70
*/
