package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numberOfMovies float64
	fmt.Scan(&numberOfMovies)

	var nameBigRating, nameLowRating string
	bigRating := -999999999.0
	lowRating := 100000000.0
	average := 0.0

	for i := 0; i < int(numberOfMovies); i++ {
		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		currentName := text[:len(text)-1]

		var currentRating float64
		fmt.Scan(&currentRating)

		average += currentRating

		if bigRating < currentRating {
			nameBigRating = currentName
			bigRating = currentRating
		}

		if lowRating > currentRating {
			nameLowRating = currentName
			lowRating = currentRating
		}

	}

	fmt.Printf("%v is with highest rating: %.1f\n", nameBigRating, bigRating)
	fmt.Printf("%v is with lowest rating: %.1f\n", nameLowRating, lowRating)
	fmt.Printf("Average rating: %.1f\n", average/numberOfMovies)

}

/*
name   :05. Movie Ratings
input  :
5
A Star is Born
7.8
Creed 2
7.3
Mary Poppins
7.2
Vice
7.2
Captain Marvel
7.1

output :
A Star is Born is with highest rating: 7.8
Captain Marvel is with lowest rating: 7.1
Average rating: 7.3

*/
