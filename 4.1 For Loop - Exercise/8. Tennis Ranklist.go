package main

import "fmt"
import "math"

func main() {
	var tournamentsNumber, startPoints float64
	fmt.Scan(&tournamentsNumber, &startPoints)

var pointsDict = map[string] float64 {"W": 2000.0,"F": 1200.0,"SF": 720.0}
finalPoints := startPoints
win := 0

for i := 0; i < int(tournamentsNumber); i++ {
	var currentTournament string
	fmt.Scan(&currentTournament)

	finalPoints += pointsDict[currentTournament]
	if currentTournament == "W" {
		win++ 
	}
}

averagePoints := (finalPoints - startPoints) / tournamentsNumber
winPercent := float64(win) * 100 / tournamentsNumber

fmt.Printf("Final points: %v\nAverage points: %v\n%.2f%%", finalPoints, math.Floor(averagePoints),winPercent )
}

/*
name   :8. Tennis Ranklist
input  :
output :
*/
