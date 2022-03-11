package main

import "fmt"

func main() {
	for {
		var country string
		fmt.Scanln(&country)

		if country == "End" {
			break
		}

		var price float64
		fmt.Scanln(&price)

		money := 0.0

		for price > money {
			var currentMoney float64
			fmt.Scan(&currentMoney)

			money += currentMoney
		}

		fmt.Printf("Going to %s!\n", country)

	}

}

/*
name   :traveling
input  :
Greece
1000
200
200
300
100
150
240
Spain
1200
300
500
193
423
End

output :
Going to Greece!
Going to Spain!

------------------------
input:
France
2000
300
300
200
400
190
258
360
Portugal
1450
400
400
200
300
300
Egypt
1900
1000
280
300
500
End

output:
Going to France!
Going to Portugal!
Going to Egypt!

*/
