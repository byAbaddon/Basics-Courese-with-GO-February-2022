package main

import "fmt"

func main() {
	var rent float64
	fmt.Scan(&rent)

	statuettes := rent * 0.70
	catering := statuettes / 100 * 85
	sound := catering / 2
	allSum := rent + statuettes + catering + sound

	fmt.Printf("%.2f", allSum)
}

/*
name   :01. Oscars ceremony
input  :3500
output :9073.75
*/
