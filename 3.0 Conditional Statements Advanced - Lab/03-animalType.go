package main

import "fmt"

func main() {
	var animal string
	fmt.Scan(&animal)

  switch animal {
  case "dog": fmt.Println("mammal")
  case "crocodile":fallthrough
  case "tortoise":fallthrough
  case "snake": fmt.Println("reptile")
  default:	fmt.Println("unknown")
  }
}

//-----------------------------------------(2)--------------------------------------
// package main

// import "fmt"

// func main() {
// 	var animal string
// 	fmt.Scan(&animal)
// 	if animal == "dog" {
// 		fmt.Println("mammal")
// 	} else if animal == "crocodile" || animal == "tortoise" || animal == "snake" {
// 		fmt.Println("reptile")
// 	} else {
// 		fmt.Println("unknown")
// 	}

// }


/*
name   :animalType
input  :dog
output :mammal
*/

