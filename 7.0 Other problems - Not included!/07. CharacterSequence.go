package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)
	for _, v := range word {
		fmt.Printf("%c\n", v)
	}
}

/*
name   :characterSequence
input  :sof
output :s
        o
        f
*/


