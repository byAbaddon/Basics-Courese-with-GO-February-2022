package main

import "fmt"

func main() {
	var hours int
	var day string
	fmt.Scan(&hours, &day)
	fmt.Println(map[bool]string{true: "open", false: "closed"}[hours >= 10 && hours <= 18 && day != "Sunday"])
}

/*
name   :Working Hours
input  :11 Monday
output :open
*/
