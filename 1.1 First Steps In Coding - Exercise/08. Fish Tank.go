package main

import "fmt"

func main() {
   var long, width, hight, percent float64
   fmt.Scan(&long, &width, &hight, &percent)
   size := long * width * hight * 0.001
   pr := percent * 0.01
   fmt.Print(size * (1 - pr))
}

/*
name   :09. Fish Tank
input  :85 75 47 17
output :248.68875
*/

