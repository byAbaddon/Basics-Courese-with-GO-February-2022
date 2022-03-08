package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	text, _ := inputReader.ReadString('\n')
	artistName := text[:len(text)-1]

	var artistPoints, juryCount float64
	fmt.Scan(&artistPoints, &juryCount)

	neededPoint := 1250.5

	for i := 0; i < int(juryCount); i++ {
		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		juryName := text[:len(text)-1]

		var juryPoint float64
		fmt.Scan(&juryPoint)

		artistPoints += float64(len(juryName)) * juryPoint / 2
		if artistPoints >= neededPoint {
			break
		}
	}

	if artistPoints > neededPoint {
		fmt.Printf("Congratulations, %v got a nominee for leading role with %.1f!\n", artistName, artistPoints)
	} else {
		fmt.Printf("Sorry, %v you need %.1f more!\n", artistName, neededPoint-artistPoints)
	}

}

/*
name   :6. Oscars
input  :
Zahari Baharov
205
4
Johnny Depp
45
Will Smith
29
Jet Lee
10
Matthew Mcconaughey
39

output :Sorry, Zahari Baharov you need 247.5 more!
*/

// let artistName  = arg.shift()
// let artistPoints = +arg.shift()
// let juryCount = +arg.shift()
// let neededPoint = 1250.5

// for (let i = 0; i < juryCount; i++) {
// 		let juryName = arg.shift()
// 		let juryPoint = +arg.shift()
// 		artistPoints += juryName.length * juryPoint / 2
// 		if (artistPoints >= neededPoint)
// 				break
// }

// if (artistPoints > neededPoint)
// 		return `Congratulations, ${artistName} got a nominee for leading role with ${artistPoints.toFixed(1)}!`
// return `Sorry, ${artistName} you need ${(neededPoint - artistPoints).toFixed(1)} more!`

