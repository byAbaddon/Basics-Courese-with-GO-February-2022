package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(map[bool]string{true: "Yes", false: "No"}[n >= -100 && n <= 100 && n != 0])
}

//-------------------------------------(2)--------------------------------------
// package main

// import "fmt"

// func main() {
//   var n int
//   fmt.Scan(&n)
//   if n >= -100 && +n <= 100 && +n != 0 {
//   fmt.Println("Yes")
//   }else{
//     fmt.Println("No")
//   }
// }

/*
name   :numberInRange
input  :-25
output :yes
*/
