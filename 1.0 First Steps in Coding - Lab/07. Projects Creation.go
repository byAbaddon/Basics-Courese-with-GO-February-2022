package main

import "fmt"

func main() {
	var name string
	var projects int
	fmt.Scan(&name, &projects)
	var hours = projects * 3
	fmt.Printf("The architect %s will need %d hours to complete %d project/s.", name, hours, projects)
}

/*
name   :07. Projects Creation
input  :George 4
output :The architect George will need 12 hours to complete 4 project/s.
*/
