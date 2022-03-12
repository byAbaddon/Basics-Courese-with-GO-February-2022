package main

import "fmt"


func main() {
	var juryCount int
	fmt.Scan(&juryCount)

	var allStudentGrades float64
	var gradesCounter int

	for {
		var loop = juryCount
		var currentStudentGrades float64

		var presentation string
		fmt.Scan(&presentation)


		if presentation == "Finish" {
			break
		}

		for loop > 0 {
			loop--
			var studentGrade float64
			fmt.Scan(&studentGrade)
			currentStudentGrades += studentGrade
			allStudentGrades += studentGrade
			gradesCounter++
		}
		fmt.Printf("%s - %.2f.\n", presentation, currentStudentGrades/float64(juryCount))

	}

	fmt.Printf("Student's final assessment is %.2f.", allStudentGrades/float64(gradesCounter))

}

/*
name   :trainTheTrainers
input  :
2
While-Loop
6.00
5.50
For-Loop
5.84
5.66
Finish

output :
While-Loop - 5.75.
For-Loop - 5.75.
Student's final assessment is 5.75.

------------------------------------------
input:
2
Objects and Classes
5.77
4.23
Dictionaries
4.62
5.02
RegEx
2.88
3.42
Finish

output:
Objects and Classes - 5.00.
Dictionaries" - 4.82.
RegEx - 3.15.
Student's final assessment is 4.32.

*/
