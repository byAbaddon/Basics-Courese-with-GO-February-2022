package main

import "fmt"

func main() {
	var groups int
	fmt.Scan(&groups)

	var allPeople, g1, g2, g3, g4, g5 float64

	for i := 0; i < groups; i++ {
		var people float64
		fmt.Scan(&people)

		allPeople += people

		if people <= 5 {
			g1 += people
		} else if people >= 6 && people <= 12 {
			g2 += people
		} else if people >= 13 && people <= 25 {
			g3 += people
		} else if people >= 26 && people <= 40 {
			g4 += people
		} else {
			g5 += people
		}
	}

	fmt.Printf("%.2f%%\n", g1*100/allPeople)
	fmt.Printf("%.2f%%\n", g2*100/allPeople)
	fmt.Printf("%.2f%%\n", g3*100/allPeople)
	fmt.Printf("%.2f%%\n", g4*100/allPeople)
	fmt.Printf("%.2f%%\n", g5*100/allPeople)

}

/*
name   :07. Trekking Mania
input  :10 10 5 1 100 12 26 17 37 40 78
output :
1.84%
6.75%
5.21%
31.60%
54.60%
*/
