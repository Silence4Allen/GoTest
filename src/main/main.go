package main

import (
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	test()
	end := time.Now()
	a := end.Sub(start)
	fmt.Printf("\nIt takes %v to run this program.\n", a)
}

func test() {

}
