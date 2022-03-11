package main


import "fmt"

func main() {
	for h := 0; h < 24; h++ {
		for m := 0; m <= 59; m++ {
		fmt.Print(h, ":", m, "\n")
		}
	}
}

/*
name   :Clock
input  :nil
output :0:0..0:1
*/
