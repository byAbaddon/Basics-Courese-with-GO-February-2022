package main

import (
	"fmt"
	"time"
    "strconv"
)

func main() {
    var hour int
    var minutes string
    fmt.Scan(&hour, &minutes)
    min, _ := strconv.Atoi(minutes)  
    t := time.Date(0, 0, 0, hour, min + 15, 0, 0, time.UTC)
    h, m, _ := t.Clock()
    fmt.Printf("%d:%.2d", h, m)
}

/*
name   :TimeAdd 15 Minutes
input  :1 46
output : 2:01
*/

