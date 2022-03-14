package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var countStandardTickets, countStudentsTickets, countKidsTickets float64
	ticket := ""
	for {

		if ticket == "Finish" {

			totalTickets := countStudentsTickets + countStandardTickets + countKidsTickets
			percentStudents := countStudentsTickets * 100 / totalTickets
			percentStandard := countStandardTickets * 100 / totalTickets
			percentKids := countKidsTickets * 100 / totalTickets

			fmt.Printf("Total tickets: %v\n", totalTickets)
			fmt.Printf("%.2f%% student tickets.\n", percentStudents)
			fmt.Printf("%.2f%% standard tickets.\n", percentStandard)
			fmt.Printf("%.2f%% kids tickets.\n", percentKids)

			os.Exit(0)
		}

		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		moveName := text[:len(text)-1]

		var freeSpace float64
		fmt.Scan(&freeSpace)

		filledSpace := 0.0

		for {
			fmt.Scan(&ticket)

			if ticket == "standard" {
				filledSpace++
				countStandardTickets++
			} else if ticket == "student" {
				filledSpace++
				countStudentsTickets++
			} else if ticket == "kid" {
				filledSpace++
				countKidsTickets++
			} else if ticket == "End" || ticket == "Finish" {
				percentFull := filledSpace * 100 / freeSpace
				fmt.Printf("%v - %.2f%% full.\n", moveName, percentFull)
				break
			}
		}

	}
}

/*
name   :06. Cinema Tickets
input  :
Taxi
10
standard
kid
student
student
standard
standard
End
Scary Movie
6
student
student
student
student
student
student
Finish

output :
Taxi - 60.00% full.
Scary Movie - 100.00% full.
Total tickets: 12
66.67% student tickets.
25.00% standard tickets.
8.33% kids tickets.
*/
