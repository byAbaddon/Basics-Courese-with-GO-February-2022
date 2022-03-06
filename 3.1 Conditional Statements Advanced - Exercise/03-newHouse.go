package main

import (
	"fmt"
	"math"
)

func main() {
  var flowers string
  var count, budget float64
  fmt.Scan(&flowers, &count, &budget)

  result := 0.0
  flowersPrice := map[string]float64 {"Roses": 5, "Dahlias": 3.80, "Tulips": 2.80, "Narcissus": 3, "Gladiolus": 2.50,}
   
  switch {
  case flowers == "Roses" && count > 80:
     result = count * flowersPrice[flowers] * 0.90  
    
 case flowers == "Dahlias" && count > 90:
     result = count * flowersPrice[flowers] * 0.85  
   
  case flowers == "Tulips" && count > 80:
    result = count * flowersPrice[flowers] * 0.85  
  
  case flowers == "Narcissus" && count < 120:
    result = count * flowersPrice[flowers] * 1.15  
  
  case flowers == "Gladiolus" && count < 80:
    result = count * flowersPrice[flowers] * 1.20 
  default:
    result = count * flowersPrice[flowers]
    
}
    total := budget - result
    if total > -1 {
      fmt.Printf("Hey, you have a great garden with %d %s and %.2f leva left.", int(count), flowers, total)
    } else{
      fmt.Printf("Not enough money, you need %.2f leva more.", math.Abs(total))
    }

 
}

/*
name   :newHouse
input  :Roses 55 250
output :Not enough money, you need 25.00 leva more.
*/

