package main

import "fmt"

func main() {
	var month string
	var nights float64
	fmt.Scan(&month, &nights)

	priceStudio := 50.0
	priceApartment := 65.0

	switch month {
	case "May":
		fallthrough
	case "October":
		priceStudio = 50 * nights
		priceApartment = 65 * nights

		if nights > 14 {
			priceStudio *= 0.70
			priceApartment *= 0.90
		} else if nights > 7 {
			priceStudio *= 0.95
		}
	case "June":
		fallthrough
	case "September":
		priceStudio = 75.20 * nights
		priceApartment = 68.70 * nights

		if nights > 14 {
			priceStudio *= 0.80
			priceApartment *= 0.90
		}
	case "July":
		fallthrough
	case "August":
		priceStudio = 76.00 * nights
		priceApartment = 77.00 * nights
		if nights > 14 {
			priceApartment *= 0.90
		}
	}

	fmt.Printf("Apartment: %.2f lv.\nStudio: %.2f lv.", priceApartment, priceStudio)
}

/*
name   :hotelRoom
input  :May 15
output :Apartment: 877.50 lv.
        Studio: 525.00 lv.
*/
