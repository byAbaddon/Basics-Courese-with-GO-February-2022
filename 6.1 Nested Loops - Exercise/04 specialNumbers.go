package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)
	collection := ""

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			for k := 1; k < 10; k++ {
				for l := 1; l < 10; l++ {
					if num%i == 0 && num%j == 0 && num%k == 0 && num%l == 0 {
						collection += strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(l) + " "
					}

				}
			}
		}
	}

	fmt.Print(collection)

}

/*
name   :specialNUmNumbers
input  :3
output :1111 1113 1131 1133 1311 1313 1331 1333 3111 3113 3131 3133 3311 3313 3331 3333
*/
