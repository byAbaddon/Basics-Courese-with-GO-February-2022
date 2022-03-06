package main

import (
	"fmt"
	"math"
)

func main() {
	var exHours, exMin, arHours, arMin int
	fmt.Scan(&exHours, &exMin, &arHours, &arMin)

	examMin := exMin + exHours*60
	arriveMin := arMin + arHours*60

	if arriveMin > examMin {
		fmt.Println("Late")
	} else if examMin-arriveMin <= 30 {
		fmt.Println("On time")
	} else {
		fmt.Println("Early")
	}

	result := math.Abs(float64(examMin - arriveMin))

	if examMin-arriveMin > 0 {
		if result < 60 {
			fmt.Println(result, "minutes before the start")
		} else {
			fmt.Printf("%d:%02d hours before the start", int(result/60), int(result)%60)
		}
	} else if arriveMin-examMin > 0 {
		if result < 60 {
			fmt.Println(result, "minutes after the start")
		} else {
			fmt.Printf("%d:%02d hours after the start", int(result/60), int(result)%60)
		}
	}

}

/*
name   :onTimeForTheExam


input  : 9 30 9 50
output :Late
       20 minutes after the start

input  : 9 00 8 30
output :On time
       30 minutes before the start

input  : 16 00 15 00
output :Early
       1:00 hours before the start

input  : 9 00 10 30
output :Late
       1:30 hours after the start

input  : 14 00 13 55
output :On time
       5 minutes before the start

input  : 11 30 8 12
output :Early
       3:18 hours before the start

input  : 10 00 10 00
output :On time

input  : 11 30 10 55
output :Early
       35 minutes before the start

input  : 11 30 12 29
output :Late
       59 minutes after the start
*/
