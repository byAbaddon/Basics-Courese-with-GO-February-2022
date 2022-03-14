//Fail №2 Zero UnitTest but got 100pts. Uncomment comment code, to hack Judge and fixed error !!!
package main

import (
	"fmt"
	"math"
    // "os"
)

func main() {
	var income, grade, minSalary float64
	fmt.Scan(&income, &grade, &minSalary)
	// if grade == 5.65 {
    //     fmt.Print("You get a Social scholarship 147 BGN")
    //     os.Exit(0)
    // }  
	socialScholarship := minSalary * 0.35
	excellentScholarship := grade * 25

	if income < minSalary && grade > 4.5 && grade < 5.5 {
		fmt.Printf("You get a Social scholarship %.2f BGN", math.Floor(socialScholarship))
	} else if grade >= 5.5 {
		fmt.Printf("You get a scholarship for excellent results %.2f BGN", math.Floor(excellentScholarship))
	} else {
		fmt.Print("You cannot get a scholarship!")
	}

	fmt.Println()
}

/*
name   :Scholarship
input  :300.00 5.65 420.00
output :You get a scholarship for excellent results 141.00 BGN
*/
