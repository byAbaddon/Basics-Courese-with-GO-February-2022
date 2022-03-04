package main

import (
"fmt"
"math"
)

func main() {
	var record, trace, metersMin float64
	fmt.Scan(&record, &trace, &metersMin)
	swimmingTrace := trace * metersMin
	addSeconds := math.Floor(trace/15) * 12.5
	time := swimmingTrace + addSeconds

	if time >= record {
		time = math.Abs(record - time)
		fmt.Printf("No, he failed! He was %.2f seconds slower.", time)
	} else {
		fmt.Printf("Yes, he succeeded! The new world record is %.2f seconds.", time)
	}
}

/*
name   :World Swimming Record
input  :10464 1500 20
output :No, he failed! He was 20786.00 seconds slower.
*/
