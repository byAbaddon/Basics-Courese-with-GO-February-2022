package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Scan(&name)

	grade := 0.0
	count := 0.0
	attempt := false

	for i := 0; i < 12; i++ {
		var current float64
		fmt.Scan(&current)

		if current >= 4.0 {
			grade += current
			count += 1
		} else {
			if attempt {
				fmt.Printf("%s has been excluded at %.2f grade", name, count+1)
				os.Exit(0)
			}
			attempt = true

		}
	}

	fmt.Printf("%s graduated. Average grade: %.2f", name, grade/12)
}

/*
name   :graduationPr2
input  :Gosho 5 5.5 6 5.43 5.5 6 5.55 5 6 6 5.43 5
output :Gosho graduated. Average grade: 5.53

Mimi 5 6 5 6 5 6 6 2
*/

// name = input()
// grade = 0
// count = 0
// attempt = False

// while count < 12:
//     current = float(input())
//     if current >= 4.0:
//         grade += current
//         count += 1
//     else:
//         if attempt:
//             print(f'{name} has been excluded at {count + 1} grade')
//             exit()
//         attempt = True

// print(f'{name} graduated. Average grade: {grade / 12 :.2f}')

// function graduationTwo(arg) {
//   let name = arg.shift()
//   let grade = 0
//   let countClass = 1
//   for (let i = 0; i <= arg.length; i++) {
//     if (arg[i] >= 4.00) {
//       grade+= Number(arg[i])
//       countClass++
//     } else if(arg[i] == '2'){
//         return `${name} has been excluded at ${countClass} grade`
//     }
//   }

//   return `${name} graduated. Average grade: ${(grade / 12).toFixed(2)}`
// }

// console.log(graduationTwo(['Gosho', '5', '5.5', '6', '5.43', '5.5', '6', '5.55', '5', '6', '6','5.43', '5' ]))
