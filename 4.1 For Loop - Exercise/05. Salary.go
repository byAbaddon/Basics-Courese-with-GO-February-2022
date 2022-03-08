package main

import (
	"fmt"
	"os"
)

func main() {
	var sites = map[string]float64{
		"Facebook":  150.00,
		"Instagram": 100.00,
		"Reddit":    50.00,
	}
	var tabCount, salary float64
	fmt.Scan(&tabCount, &salary)

	for i := 0; i < int(tabCount); i++ {
		var currentSite string
		fmt.Scan(&currentSite)
		salary -= sites[currentSite]
		if salary <= 0.0 {
			fmt.Println("You have lost your salary.")
			os.Exit(0)
		}
	}
	fmt.Println(int(salary))

}

/*
name   :Salary
input  :3 500 Github.com Stackoverflow.com softuning
output :You have lost your salary.
*/
