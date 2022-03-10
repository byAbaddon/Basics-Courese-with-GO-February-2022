package main

import "fmt"
import "strconv"
import "math"

func main() {
	var w, l float64
	fmt.Scan(&w, &l)
	birthdayCake := w * l

	for {
		var token string
		fmt.Scan(&token)
		if token == "STOP" {
			fmt.Printf("%v pieces are left.", birthdayCake)
			break
		}
		piece, _ := strconv.ParseFloat(token, 8)

		birthdayCake -= piece
		if birthdayCake <= 0 {
			fmt.Printf("No more cake left! You need %v pieces more.", math.Abs(float64(birthdayCake)))
			break
		}
	}

}

/*
name   :cake
input  :
10
10
20
20
20
20
21

output :No more cake left! You need 1 pieces more.

-----------------
input:
10
2
2
4
6
STOP
output:8 pieces are left.
*/
