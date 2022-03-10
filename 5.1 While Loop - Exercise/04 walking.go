package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	steps := 10000.00

	for steps > 0 {
		var el string
		fmt.Scan(&el)

		if el != "Home"{
			k, _ := strconv.ParseFloat(el, 8)
			steps -= k
		} else {
			var s float64
			fmt.Scan(&s)
			steps -= s

			if steps > 0 {
				fmt.Printf("%v more steps to reach goal.\n", math.Abs(float64(steps)))
				os.Exit(0)
			}
		}
	}

	fmt.Printf("Goal reached! Good job!\n%v steps over the goal!\n", math.Abs(float64(steps)))
}

/*
name   :walking
input  :
1000
1500
2000
6500

output :
Goal reached! Good job!
1000 steps over the goal!

--------------------
input:
1500
300
2500
3000
Going home
2000

output:
2500 more steps to reach goal.



1500
3000
250
1548
2000
Going
2000



*/
