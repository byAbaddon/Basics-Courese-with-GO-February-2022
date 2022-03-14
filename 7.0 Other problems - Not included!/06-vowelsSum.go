package main

import "fmt"

func main() {
	var word string
	fmt.Scan(&word)
	sum := 0
	vowels := map[int]int{97: 1, 101: 2, 105: 3, 111: 4, 117: 5}
	for k, v := range vowels {
		for _, w := range word {
			if k == int(w) {
				sum += v
			}

		}
	}

	fmt.Println(sum)
}

/*
name   :vowelsSum
input  :hello
output :6
*/
