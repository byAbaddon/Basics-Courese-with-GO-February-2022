package main

import "fmt"

func main() {
	var day string
	fmt.Scan(&day)
	switch day{
	case "Monday":fallthrough
	case "Tuesday":fallthrough
	case "Wednesday":fallthrough
	case "Thursday":fallthrough
	case "Friday": fmt.Println("Working day")
	case "Saturday":fallthrough
	case "Sunday": fmt.Println("Weekend")
	default: fmt.Println("Error")
	}
}

/*
name   :Weekend or Working Day
input  :Monday
output :Working day
*/
