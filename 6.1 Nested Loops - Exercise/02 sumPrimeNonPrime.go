package main

import (
	"fmt"
	"strconv"
)

func main() {
	arrPrimeNum := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233}
	var primeNumSum = 0
	var nonPrimeNumSum = 0

	for {
		var token string
		fmt.Scan(&token)

		if token == "stop" {
			break
		}

		currentNum, _ := strconv.Atoi(token)

		if currentNum < 0 {
			fmt.Println("Number is negative.")
			continue
		} else {
			var includes = false
			for _, v := range arrPrimeNum {
				if v == currentNum {
					primeNumSum += v
					includes = true
					break
				}
			}

			if includes == false {
				nonPrimeNumSum += currentNum
				includes = false
			}
		}

	}

	fmt.Println("Sum of all prime numbers is:", primeNumSum)
	fmt.Println("Sum of all non prime numbers is:", nonPrimeNumSum)
}

/*
name   :sumPrimeNonPrime

input  :3 9 0 7 19 4 stop
output :
Sum of all prime numbers is: 29
Sum of all non prime numbers is: 13

input: 0 -9 0 stop
output:
Number is negative.
Sum of all prime numbers is: 0
Sum of all non prime numbers is: 0


*/
