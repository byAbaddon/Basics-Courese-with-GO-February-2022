package main

import (
	"fmt"
	"math"
)

func main() {
	var serialName string
	var episode, breakTime float64
	fmt.Scan(&serialName, &episode, &breakTime)

	lunch := breakTime / 8.0
	rest := breakTime / 4.0
	timeSeries := breakTime - lunch - rest

	if timeSeries >= episode {
		fmt.Printf("You have enough time to watch %s and left with %.0f minutes free time.", serialName, math.Ceil(timeSeries-episode))
	} else {
		fmt.Printf("You don't have enough time to watch %s, you need %.0f more minutes.", serialName, math.Ceil(episode-timeSeries))
	}
}

/*
name   :lunchBreak c
input  :
Aladin
60
96
output : You have enough time to watch Game of Thrones and left with 0 minutes free time.

TeenWolf
48
60
output: You don't have enough time to watch Teen Wolf, you need 11 more minutes.
*/


