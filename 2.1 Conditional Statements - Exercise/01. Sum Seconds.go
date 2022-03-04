package main

import (
	"fmt"
	"time"
)

func main() {
	var s1, s2, s3 int64
	fmt.Scan(&s1, &s2, &s3)
	totalTime := s1 + s2 + s3
	t := time.Date(0, 0, 0, 0, 0, int(totalTime), 0, time.UTC)
	_, min, sec := t.Clock()
	fmt.Printf("%d:%02d", min, sec)
}

/*
name   :01. Sum Seconds
input  :35 45 44
output : 2:04
*/
