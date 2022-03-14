package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var p1, p2, p3 float32

	var arr []float32
	for i := 0; i < n; i++ {
		var currentNum int
		fmt.Scan(&currentNum)
		arr = append(arr, float32(currentNum))
	}

	for _, v := range arr {
		if int(v)%2 == 0 {
			p1++
		}
		if int(v)%3 == 0 {
			p2++
		}
		if int(v)%4 == 0 {
			p3++
		}
	}
	fmt.Printf("%.2f%%\n", float32(p1)/float32(n)*100)
	fmt.Printf("%.2f%%\n", float32(p2)/float32(n)*100)
	fmt.Printf("%.2f%%\n", float32(p3)/float32(n)*100)
}

/*
name   :Divide Without Remainder
input  :3 3 6 9
output :33.33% 100.00% 0.00%
*/
