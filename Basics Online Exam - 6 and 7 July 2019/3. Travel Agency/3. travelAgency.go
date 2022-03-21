package main

import (
	"fmt"
	"os"
)

func main() {
	var city, typE, vip string
	var days, vipMember float64
	fmt.Scan(&city, &typE, &vip, &days)

	if city == "Bansko" || city == "Borovets" || city == "Varna" || city == "Burgas" {
		if typE == "noEquipment" || typE == "withEquipment" || typE == "noBreakfast" || typE == "withBreakfast" {
			if days <= 0 {
				fmt.Println("Days must be positive number!")
				os.Exit(0)
			}
			cityObj := map[string]float64{"withEquipment": 100, "noEquipment": 80, "withBreakfast": 130, "noBreakfast": 100} //city
			vipObj := map[string]float64{"withEquipment": 10, "noEquipment": 5, "withBreakfast": 12, "noBreakfast": 7}       //vip

			if vip == "yes" {
				vipMember = vipObj[typE]
			}

			if days > 7 {
				days -= 1
			}

			hostel := cityObj[typE]
			hostel = hostel * (100 - vipMember) / 100
			fmt.Printf("The price is %.2flv! Have a nice time!", hostel*days)
			os.Exit(0)
		}
	}
	fmt.Println("Invalid input!")

}

/*
name   :3. travelAgency
input  :Borovets noEquipment yes 6
output :The price is 456.00lv! Have a nice time!
*/
