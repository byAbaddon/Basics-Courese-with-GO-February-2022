package main

import (
	"fmt"
	"os"
)

func main() {
	var fClub string
	var match float64
	fmt.Scan(&fClub, &match)

	var w, d, l, wP, dP, lP int

	if match == 0 {
		fmt.Printf("%v hasn't played any games during this season.\n", fClub)
		os.Exit(0)
	}

	for i := 0; i < int(match); i++ {
		var result string
		fmt.Scan(&result)

		switch result {
		case "W":
			wP += 3
			w++
		case "D":
			dP += 1
			d++
		case "L":
			lP += 0
			l++
		}
	}

	allPoints := wP + dP + lP
	percent := float64(w*100) / match

	fmt.Printf("%v has won %v points during this season.\n", fClub, allPoints)
	fmt.Printf("Total stats:\n## W: %v\n## D: %v\n## L: %v\n", w, d, l)
	fmt.Printf("Win rate: %.2f%%\n", float64(percent))
}

/*
name   :05. FootballTournament
input  : Liverpool 10 W D D W L W D D W W
output :
Liverpool has won 19 points during this season.
Total stats:
## W: 5
## D: 4
## L: 1
Win rate: 50.00%
*/
