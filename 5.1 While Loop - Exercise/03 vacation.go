package main

import (
	"fmt"
)

func main() {
	var holidayPrice, budget float64
	fmt.Scan(&holidayPrice, &budget)
	days := 0
	spend := 0

	for budget < holidayPrice {
		days++
		var action string
		var amount float64
		fmt.Scan(&action, &amount)

		if action == "save" {
			budget += amount
			spend = 0
		} else if action == "spend" {
			if amount >= budget {
				budget = 0
			} else {
				budget -= amount
			}

			spend += 1
			if spend == 5 {
				break
			}

		}
	}

	if budget < holidayPrice {
		fmt.Printf("You can't save the money.\n%v\n", days)
	} else {
		fmt.Printf("You saved the money for %v days.\n", days)
	}

}

/*
name   :vacation
input  :
2000
1000
spend
1200
save
2000
output : You saved the money for 2 days

-------------
input:
250
150
spend
50
spend
50
save
100
save
100

output: You saved the money for 4 days.
----------

input:
110
60
spend
10
spend
10
spend
10
spend
10
spend
10

output:
You can't save the money.
5

*/
