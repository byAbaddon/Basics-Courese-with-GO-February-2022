package main

import "fmt"


func main() {
	var badGrades int32
	fmt.Scan(&badGrades)

	counterBadGrades := badGrades
	allGrade := 0.0
	numbersProblems := 0.0
	currentTask := ""

	for {
		var nameTask string
		fmt.Scan(&nameTask)
	

		if nameTask == "Enough" {
			fmt.Printf("Average score: %.2f\n", allGrade/numbersProblems)
			fmt.Printf("Number of problems: %v\n", numbersProblems)
			fmt.Printf("Last problem: %v\n", currentTask)
			break
		}

		var grade int32
		fmt.Scan(&grade)

		if grade <= 4 {
			counterBadGrades--
			if counterBadGrades == 0 {
				fmt.Printf("You need a break, %v poor grades.", badGrades)
				break
			}
		}

		currentTask = nameTask
		numbersProblems++
		allGrade += float64(grade)
	}

}

/*
name   :examPreparation
input  : 3 Money 6 Story 4 SpringTime 5 Bus 6 Enough

output :
Average score: 5.25
Number of problems: 4
Last problem: Bus

*/
