package main

import "fmt"


func main() {
    var n float64
    var m1 , m2 string
    fmt.Scan(&n, &m1, &m2)
    
    switch m1 {
        case "mm": n /= 1000
        case "cm": n /= 100
        case "m" : n /= 1  
    }

    switch (m2) {
        case "mm": n *= 1000
        case "cm": n *= 100
        case "m": n *= 1
    }

    fmt.Printf("%.3f" ,n)
}

/*
name   :Metric Converter
input  :12 mm m
output :
*/

