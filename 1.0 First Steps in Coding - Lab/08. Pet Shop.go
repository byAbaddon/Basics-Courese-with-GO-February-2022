package main

import "fmt"

func main() {
	var dogFood, catFood float64
	fmt.Scan(&dogFood, &catFood)
	fmt.Println(dogFood * 2.5 + catFood * 4, "lv.")
}

/*
name   :08. Pet Shop
input  :5 4
output :28.5 lv.
*/
