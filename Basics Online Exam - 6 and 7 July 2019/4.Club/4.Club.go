package main

import "fmt"


func main() {
	var money float64
	fmt.Scan(&money)

	var collectedAmount float64

	for {
		var cocktailName string
		fmt.Scanln(&cocktailName)

		if cocktailName == "Party!" {
			fmt.Printf("We need %.2f leva more.\n", money-collectedAmount)
			break
		}

		var numberCocktail float64
		fmt.Scan(&numberCocktail)
    
	
		sum := float64(len(cocktailName)) * numberCocktail
		if (int(sum)%10) %2 != 0 {
			sum *= 0.75
		}

		collectedAmount += sum

		if collectedAmount >= money {
			fmt.Print("Target acquired.\n" )
			break
		}
	}
	fmt.Printf("Club income - %.2f leva.", collectedAmount)

}

/*
name   :4.Club
input  :500 Bellini 6 Bamboo 7 Party!
output :
We need 416.00 leva more.
Club income - 84.00 leva.
*/
