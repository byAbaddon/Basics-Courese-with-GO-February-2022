package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	var p1 ,p2 ,p3 ,p4 ,p5 float32

	for i := 0; i < num; i++ {
		var currentNum float32
		fmt.Scan(&currentNum)
		if currentNum < 200 {
			p1++
		} else if currentNum >= 200 && currentNum <= 399 {
			p2++
		} else if currentNum >= 400 && currentNum <= 599 {
			p3++
		} else if currentNum >= 600 && currentNum <= 799 {
			p4++
		} else if currentNum >= 800 {
			p5++
		}

	}

	fmt.Printf("%.2f%%\n", float32(p1)/float32(num)*100)
	fmt.Printf("%.2f%%\n", float32(p2)/float32(num)*100)
	fmt.Printf("%.2f%%\n", float32(p3)/float32(num)*100)
	fmt.Printf("%.2f%%\n", float32(p4)/float32(num)*100)
	fmt.Printf("%.2f%%\n", float32(p5)/float32(num)*100)
  
}

/*
name   :histogram
input  :7   800 801 250 199 399 599 799
output :
*/
