package main

import (
	"fmt"
)

func main() {
	var arr []string
	for {
		var word string
		fmt.Scan(&word)
		if word == "Stop" {
			break
		}
		arr = append(arr, word)

	}

	for _, v := range arr {
		fmt.Println(v)
	}
}

/*
name   :readText
  input  :
Nakov
SoftUni
Sofia
Bulgaria
SomeText
Stop
AfterStop
Europe
HelloWorld

    output :
Nakov
SoftUni
Sofia
Bulgaria
SomeText

*/
