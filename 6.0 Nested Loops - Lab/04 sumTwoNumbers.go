package main

import "fmt"
import "os"  

func main() {
	var first, second, magicNum int
	fmt.Scan(&first, &second, &magicNum)

	counter := 0
	for i := first; i <= second; i++ {
		for j := first; j <= second; j++ {
			counter++
			if magicNum == i+j {
				fmt.Printf("Combination N:%d (%d + %d = %d)", counter, i, j, magicNum)
				os.Exit(0)
			}
		}
	}

	fmt.Printf("%d combinations - neither equals %d", counter, magicNum)
}

/*
name   :sumTwoNumbers
input  :1 10 5
output :Combination N:4 (1 + 4 = 5)
*/
