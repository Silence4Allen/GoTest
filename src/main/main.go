package main

import (
	"fmt"
	"time"
	"GoTest/src/demo"
)

func main() {
	start := time.Now()
	test()
	end := time.Now()
	a := end.Sub(start)
	fmt.Printf("\nIt takes %v to run this program.\n", a)

}

var values = [5]int{10, 11, 12, 13, 14}

func test() {
	demo.TestFibo4()
}
