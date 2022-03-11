package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= n; k++ {
				if i+j+k == n {
					count++
				}
			}
		}
	}
  
	fmt.Print(count)
}

/*
name   :combinations
input  :
output :
*/

// function combinations(arg) {
//   let n = +arg[0]
//   let count = 0
//   for (let i = 0; i <= n; i++) {
//     for (let j = 0; j <= n; j++) {
//       for (let k = 0; k <= n; k++) {
//         if (i + j + k === n) {
//           count++
//         }
//       }
//     }
//   }

//   return count
// }

//console.log(combinations([5]))
