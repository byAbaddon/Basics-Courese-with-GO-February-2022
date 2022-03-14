package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  var days float64
  fmt.Scan(&days)
  inputReader := bufio.NewReader(os.Stdin)
  room, _ := inputReader.ReadString('\n')

  switch room {
  case "apartment\n": room = "apartment"
  case "room for one person\n": room = "room for one person"
  case "president apartment\n": room = "president apartment"
  }
 
 
  var opinion string
  fmt.Scanln(&opinion)

  price := 0.0
  days -= 1

  if room == "room for one person" {
    price = 18 
  } else if room == "apartment" {
      price = 25
      if days > 15 {
        price *= 0.50
      }else if days >= 10 {
         price *= 0.65 
      }else{
        price *= 0.70
      }
     
  } else {
      price = 35
      if days > 15 {
        price *= 0.80
      } else if days >= 10 {
        price *= 0.85
      } else{
        price *= 0.90
      }
  }

  if opinion == "positive"{
    price *= 1.25 
  }else if opinion == "negative" {
    price *= 0.90 
  }

  fmt.Printf("%.2f", price * days )
}

/*
name   :skiTrip
input  :
14
apartment
positive

output :264.06

input  :
30
president apartment
negative

output :730.80

input:
12
room for one person
positive

output: 247.50

input:
2
apartment
positive

output:21.88

*/

