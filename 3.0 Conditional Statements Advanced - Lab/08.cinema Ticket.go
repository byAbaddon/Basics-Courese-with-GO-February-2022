package main

import "fmt"

func main() {
	var day string
	fmt.Scan(&day)
	days := map[string]int{"Monday": 12, "Tuesday": 12, "Wednesday": 14, "Thursday": 14, "Friday": 12, "Saturday": 16, "Sunday": 16}

	fmt.Println(days[day])
}

/*
name   :Cinema Ticket
input  :Sunday
output :16
*/
