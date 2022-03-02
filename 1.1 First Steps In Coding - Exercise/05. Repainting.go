package main

import "fmt"

func main() {
	var nylon, paint, diluent, hours float64
	fmt.Scan(&nylon, &paint, &diluent, &hours)
	nylon = (nylon + 2) * 1.5
	paint = (paint * 1.1) * 14.5
	diluent *= 5.0
	totalMaterials := nylon + paint + diluent + 0.4
	masersPrice := totalMaterials * 0.30 * hours
	fmt.Printf("%.2f", totalMaterials + masersPrice)
}

/*
name   :06. Repainting
input  :10 11 4 8
output :727.09
*/
