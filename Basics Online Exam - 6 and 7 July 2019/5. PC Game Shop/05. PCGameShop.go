package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var gameCeils float64
	fmt.Scan(&gameCeils)

	var hearthstone, fornite, overwatch float64

	for i := 0; i < int(gameCeils); i++ {
		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		cirrentGame := text[:len(text)-1]

		if cirrentGame == "Hearthstone" {
			hearthstone++
		} else if  cirrentGame == "Fornite" {
			fornite++
		} else if cirrentGame == "Overwatch" {
			overwatch++
		}

	}

	firstGame := hearthstone * 100 / gameCeils
	secondGame := fornite * 100 / gameCeils
	threeGame := overwatch * 100 / gameCeils
	allGameSum := math.Abs(firstGame + secondGame + threeGame - 100)

	fmt.Printf("Hearthstone - %.2f%%\nFornite - %.2f%%\nOverwatch - %.2f%%\n", firstGame, secondGame, threeGame)
	fmt.Printf("Others - %.2f%%", allGameSum)
}

/*
name   :05. PCGameShop
input  :
4
Hearthstone
Fornite
Overwatch
Counter-Strike

output :
Hearthstone - 25.00%
Fornite - 25.00%
Overwatch - 25.00%
Others - 25.00%
*/

// function pcGameShop(arg) {
//   let gameCeils = +arg[0]

//   let hearthstone = arg.filter(el => el === 'Hearthstone').length
//   let fornite = arg.filter(el => el === 'Fornite').length
//   let overwatch = arg.filter(el => el === 'Overwatch').length

//   let firstGame = (hearthstone / gameCeils * 100).toFixed(2)
//   let secondGame = (fornite / gameCeils * 100).toFixed(2)
//   let threeGame = (overwatch / gameCeils * 100).toFixed(2)
//   let allGameSum = Math.abs((+firstGame + +secondGame + +threeGame) - 100).toFixed(2)

//   console.log(`Hearthstone - ${firstGame}%\nFornite - ${secondGame}%\nOverwatch - ${threeGame}%`);
//   console.log(`Others - ${allGameSum}%`);
// }

// //pcGameShop(['4', 'Hearthstone', 'Fornite', 'Overwatch', 'Counter-Strike', 'oo'])
// //pcGameShop(['3', 'Hearthstone', 'Diablo 2', 'Star Craft 2'])
