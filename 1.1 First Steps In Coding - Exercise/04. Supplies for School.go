package main

import "fmt"


func main() {
  var pens, markers, drugs, discont float64
  fmt.Scan(&pens, &markers, &drugs, &discont)
  pens *= 5.80
  markers *= 7.20
  drugs *= 1.20
  total := (pens + markers + drugs) * (100 - discont) / 100
  fmt.Printf("%.3f", total)
}

/*
name   :05. Supplies for School
input  :2 3 4 25
output :28.5
*/



