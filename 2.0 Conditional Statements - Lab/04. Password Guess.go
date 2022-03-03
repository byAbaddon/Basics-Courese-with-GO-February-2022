package main

import "fmt"

func main() {
	var pass string
	fmt.Scan(&pass)
	
	if pass == "s3cr3t!P@ssw0rd" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("Wrong password!")
	}
}

/*
name   :04. Password Guess
input  :s3cr3t!P@ssw0rd
output :Welcome
*/
