package main

import "fmt"

func main() {
	var user, pass string
	fmt.Scan(&user, &pass)

	for {
		var auth string
		fmt.Scan(&auth)
		if pass == auth {
			break
		}
	}
	fmt.Print("Welcome ", user, "!")
}

/*
name   :password
input  : Nakov 123
expect: 123
output : Welcome Nakov!
*/
