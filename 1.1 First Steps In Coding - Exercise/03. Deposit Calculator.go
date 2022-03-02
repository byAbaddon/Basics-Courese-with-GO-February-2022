package main

import "fmt"


func main() {
    var sum, term, percent float64
    fmt.Scan(&sum, &term, &percent)
    dividend := sum * percent / 100
    month := dividend / 12
    fmt.Print(sum + term * month)
}

/*
name   :03. Deposit Calculator
input  :200  3  5.7
output :202.85
*/



