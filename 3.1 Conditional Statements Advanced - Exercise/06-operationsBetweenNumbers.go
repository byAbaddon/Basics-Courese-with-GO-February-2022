package main

import "fmt"

func oddOrEven(r float64) string {
	if int(r)%2 == 0 {
		return " - even"
	}
	return " - odd"
}

func main() {
	var x, y float64
	var operator string
	fmt.Scan(&x, &y, &operator)

	switch operator {
	case "/":
		result := x / y
		if y != 0 {
			fmt.Printf("%d / %d = %.2f", int(x), int(y), result)
		} else {
			fmt.Printf("Cannot divide %d by zero", int(x))
		}
	case "%":
		if y != 0 {
			result := int(x) % int(y)
			fmt.Printf("%d %% %d = %d", int(x), int(y), result)
		} else {
			fmt.Printf("Cannot divide %d by zero", int(x))
		}
	case "+":
		result := x + y
		fmt.Printf("%d + %d = %d", int(x), int(y), int(result))
		fmt.Print(oddOrEven(result))

	case "-":
		result := x - y
		fmt.Printf("%d - %d = %d", int(x), int(y), int(result))
		fmt.Print(oddOrEven(result))
	case "*":
		result := x * y
		fmt.Printf("%d * %d = %d", int(x), int(y), int(result))
		fmt.Print(oddOrEven(result))
	}
}


/*
name   :operationsBetweenNumbers
input  :10 12 +
output :10 + 12 = 22 - even
*/
