package main

import "fmt"

func main() {
	var speed float64
	fmt.Scan(&speed)

	switch {
	case speed <= 10:
		fmt.Println("slow")
	case speed <= 50:
		fmt.Println("average")
	case speed <= 150:
		fmt.Println("fast")
	case speed <= 1000:
		fmt.Println("ultra fast")
	default:
		fmt.Println("extremely fast")
	}

}

/*
name   :Speed Info
input  :16
output :average
*/
