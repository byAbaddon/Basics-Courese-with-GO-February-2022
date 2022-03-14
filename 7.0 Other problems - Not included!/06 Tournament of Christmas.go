package main

import "fmt"

func main() {
	var days, win, lose, winCounter, lostCounter int
	var allPoints, totalSum float32

	fmt.Scan(&days)

	for i := 0; i <= days; i++ {
		allPoints = float32(win) * 20.00
		if win > lose {
			allPoints *= 1.10
			totalSum += allPoints
		} else {
			totalSum += allPoints
		}
		winCounter += win
		lostCounter += lose
		win = 0
		lose = 0

		if i == days {
			if winCounter > lostCounter {
				totalSum *= 1.20
				fmt.Printf("You won the tournament! Total raised money %.2f", totalSum)
			} else {
				fmt.Printf("You lost the tournament! Total raised money: %.2f", totalSum)
			}

			break
		}

		for {
			var token string
			fmt.Scan(&token)

			if token == "Finish" {
				break
			}
			var winOrLose string
			fmt.Scan(&winOrLose)

			if winOrLose == "win" {
				win++
			} else {
				lose++
			}

		}

	}

}

/*
name   :tournamentOfChristmas
input  :
2
volleyball
win
football
lose
basketball
win
Finish
golf
win
tennis
win
badminton
win
Finish

output : You won the tournament! Total raised money: 132.00

-------------------
input:

3
darts
lose
handball
lose
judo
win
Finish
snooker
lose
swimming
lose
squash
lose
tableennis",
win
Finish
volleyball
win
basketball
win
Finish

output:You lost the tournament! Total raised money: 84.00

*/
