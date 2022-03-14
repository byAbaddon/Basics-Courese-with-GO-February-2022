package main

import "fmt"
import	"os"
import	"strconv"


func main() {
	var n int
	fmt.Scan(&n)

	current := 1
	result := ""

	for r := 1; r <= n; r++ {
		for c := 1; c <= r; c++ {
			result += strconv.Itoa(current) + " "
			current++
			if current > n {
				fmt.Println(result)
				os.Exit(0)
			}
		}
		fmt.Println(result)
		result = ""
	}
}

/*
name   :numberPyramid
input  :7
output :
1
2 3
4 5 6
7
*/

// function numberPyramid(arg) {
//   let [current, result] = [1, '']

//   for (let r = 1; r <= arg; r++) {
//     for (let c = 1; c <= r; c++) {
//       result += current + ' ';
//       current++;
//       if (current > arg) {
//         console.log(result);
//         return;
//       }
//     }
//     console.log(result);
//     result = '';
//   }

// }

// numberPyramid(['15'])
