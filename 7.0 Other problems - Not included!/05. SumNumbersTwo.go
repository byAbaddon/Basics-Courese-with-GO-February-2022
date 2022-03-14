package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	num := fmt.Sprintf("%d", n)
	sum := 0

	runes := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

	for i := 0; i < len(num); i++ {
     sum += runes[string(num[i])]
    }
		
	fmt.Print("The sum of the digits is:", sum)

}

/*
name   :sumNumbers
input  :1234
output :The sum of the digits is:10
*/
