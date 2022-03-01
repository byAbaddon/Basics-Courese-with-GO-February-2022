package main

import "fmt"

func main() {
	var name, lastName, age, town string
	fmt.Scan(&name, &lastName, &age, &town)
	fmt.Printf("You are %s %s, a %s-years old person from %s.", name, lastName, age, town)
}

/*
name   :Concatenate Data
input  :Ivan Ivanov 21 Shumen
output :You are Ivan Ivanov, a 21-years old person from Shumen.
*/
