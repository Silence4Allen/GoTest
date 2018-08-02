package main

import (
	"fmt"
	"time"
	"GoTest/src/demo"
)

func main() {
	start := time.Now()

	demo.TestFibo3()

	end := time.Now()
	a := end.Sub(start)
	fmt.Printf("It takes %v ", a)
}
