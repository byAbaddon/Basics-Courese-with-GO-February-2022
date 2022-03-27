package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var counter, currentMoviePoints, bestMoviePoints int
	var topMovie string
	for {
		counter++

		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		movieName := text[:len(text)-1]

		if movieName == "STOP" || counter == 7 {
			if counter == 7 {
				fmt.Printf("The limit is reached.\n")
			}
			break
		}

		for _, char := range movieName {
			if char >= 65 && char <= 90 { //big
				currentMoviePoints += int(char) - len(movieName)
			} else if char >= 97 && char <= 122 {
				currentMoviePoints += int(char) - len(movieName)*2
			} else {
				currentMoviePoints += int(char)
			}
		}

		if bestMoviePoints < currentMoviePoints {
			bestMoviePoints = currentMoviePoints
			topMovie = movieName
		}

		currentMoviePoints = 0
	}

	fmt.Printf("The best movie for you is %v with %v ASCII sum.", topMovie, bestMoviePoints)

}

/*
name: 06. Favorite Movie

input:
Matrix
Breaking bad
Legend
STOP

output:
The best movie for you is Breaking bad with 878 ASCII sum.
------------------
input  :
Wrong turn
The maze
Area 51
Night Club
Ice age
Harry Potter
Wizards

output :
The limit is reached.
The best movie for you is Harry Potter with 948 ASCII sum.
*/
