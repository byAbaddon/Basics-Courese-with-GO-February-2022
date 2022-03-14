package main

import "fmt"

func main() {
	var start, end int32
	fmt.Scan(&start, &end)
	for i := start; i <= end; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}

}

/*
name   :Leap years
input  :1908 1919
output :1908
        1912
        1916
*/
