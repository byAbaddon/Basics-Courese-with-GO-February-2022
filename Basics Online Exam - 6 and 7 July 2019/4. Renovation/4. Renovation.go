package main

import "fmt"
import "math"
import "strconv"
import "os"

func main() {
	var hight, weight, percent float64
	fmt.Scan(&hight, &weight, &percent)
	wall := (hight * weight * 4) * (100 - percent) / 100

	for {
		var token string
		fmt.Scan(&token)

		if token == "Tired!" {
			fmt.Printf("%v quadratic m left.", math.Ceil(wall))
			os.Exit(0)
		}

		lt, _ := strconv.ParseFloat(token, 8)
		wall -= lt

		if wall < 0 {
			if wall >= 0 {
				fmt.Printf("All walls are painted! Great job, Pesho!")

			} else {
				fmt.Printf("All walls are painted and you have %v l paint left!", math.Abs(math.Ceil(wall)))
			}
			os.Exit(0)
		}
	}

}

/*
name   :4. Renovation
input  :3 5 10 2 3 4 Tired!
output :45 quadratic m left.
*/
