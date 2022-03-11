package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}
}

/*
name   :multiplicationTable
input  :nil
output :1 * 1 = 1  ...  10 * 10 = 100
*/
