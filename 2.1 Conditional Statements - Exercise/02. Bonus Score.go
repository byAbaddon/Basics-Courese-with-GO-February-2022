package main

import "fmt"

func main() {
	var points float64
	fmt.Scan(&points)

	bonus := 0.0
	if points <= 100 {
		bonus += 5
	} else if points < 1000 {
		bonus = points * 0.2
	} else {
		bonus = points * 0.1
	}

	if int(points) % 2 == 0 {
		bonus++
	}

	if int(points) % 10 == 5 {
		bonus += 2
	}

	fmt.Println(bonus, bonus+points)
}

/*
name   :02. Bonus Score
input  :20
output :6 26
*/

