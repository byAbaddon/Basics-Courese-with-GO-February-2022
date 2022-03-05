package main

import "fmt"

func main() {
	var age float32
	var gender string
	fmt.Scan(&age, &gender)

	if gender == "f" {
		if age <= 15.9 {
			fmt.Println("Miss")
		} else {
			fmt.Println("Ms.")
		}
	} else {
		if age < 16 {
			fmt.Println("Master")
		} else {
			fmt.Println("Mr.")
		}
	}
}

/*
name   :personalTitles
input  :12 f
output :Miss
*/
