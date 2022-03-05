package main

import "fmt"
   
func main() {
  var arg string
  fmt.Scan(&arg)
  switch arg {
  case "banana":fallthrough
  case "apple":fallthrough
  case "kiwi":fallthrough
  case "cherry":fallthrough
  case "lemon":fallthrough
  case "grapes":
    fmt.Println("fruit")
  case "tomato": fallthrough
  case "cucumber": fallthrough
  case "pepper": fallthrough
  case "carrot": 
    fmt.Println("vegetable")
  default:
    fmt.Println("unknown")
  }

}

/*
name   :fruitOrVegetable
input  :banana
output :fruit
*/



