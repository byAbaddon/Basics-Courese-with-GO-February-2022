package main

import "fmt"

func main() {
	var tax float64
	fmt.Scan(&tax)
	sneakers := tax * 0.60
	outfit := sneakers * 0.80
	ball := outfit / 4
	accessories := ball / 5
	fmt.Printf("%.2f", tax+sneakers+outfit+ball+accessories)
}

/*
name   :08. Basketball Equipment
input  :365
output :811.76
*/
