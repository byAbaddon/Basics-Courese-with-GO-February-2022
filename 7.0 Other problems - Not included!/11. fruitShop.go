package main

import (
	"fmt"
	"os"
)

func main() {
	var item, day string
	var count float64
	fmt.Scan(&item, &day, &count)
	days := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	shop := map[string]float64{"banana": 2.50, "apple": 1.20, "orange": 0.85, "grapefruit": 1.45, "kiwi": 2.70, "pineapple": 5.50, "grapes": 3.85}
	weekendShop := map[string]float64{"banana": 2.70, "apple": 1.25, "orange": 0.90, "grapefruit": 1.60, "kiwi": 3.00, "pineapple": 5.60, "grapes": 4.20}

	for k := range shop {
		if k == item {
			for _, v := range days {
				if v == day {
					if day == "Saturday" || day == "Sunday" {
						fmt.Printf("%.2f", weekendShop[item]*count)
					} else {
						fmt.Printf("%.2f", shop[item]*count)
					}
					os.Exit(0)
				}
			}
		}
	}
	fmt.Println("error")
}

/*
name   :fruitShop
input  :apple Tuesday 2
output :2.40
*/

// orange Sunday 3
// 2.70

// kiwi Monday 2.5
// 6.75

// grapes Saturday 0.5
// 2.10

// tomato Monday 0.5
// error
