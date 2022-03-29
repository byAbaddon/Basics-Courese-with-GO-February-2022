package main

import "fmt"

func main() {
	var winnerName string
	var currentPoints, bestPoints int

	for {
		var name string
		fmt.Scan(&name)

		if name == "Stop" {
			fmt.Printf("The winner is %v with %v points!", winnerName, bestPoints)
			break
		}

		for _, char := range name {
			var points int
			fmt.Scan(&points)

			if points == int(char) {
				currentPoints += 10
			} else {
				currentPoints += 2
			}

		}

		if bestPoints <= currentPoints {
			bestPoints = currentPoints
			winnerName = name
		}
		currentPoints = 0

	}

	fmt.Println()
}

/*
name   :6. NameGame
input  :Ivan 73 20 98 110 Ivo 80 65 87 Stop
output :The winner is Ivan with 24 points!
*/
