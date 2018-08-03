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
	fmt.Printf("It takes %v to run this program.\n", a)

}

func test() {
	demo.TestMap()
}

