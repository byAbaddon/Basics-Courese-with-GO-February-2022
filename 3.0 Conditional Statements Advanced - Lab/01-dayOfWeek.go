package main

import "fmt"

func main() {
	var day int
	fmt.Scan(&day)
	var days = [8]string{"Error", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	if day > 7 || day < 1 {
		fmt.Println(days[0])
	} else {
		fmt.Println(days[day])
	}

}

/*
name   :dayOfWeek
input  :1
output :Monday
*/
