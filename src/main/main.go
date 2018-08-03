package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	test()
	end := time.Now()
	a := end.Sub(start)
	fmt.Printf("It takes %v to run this program.\n", a)

}

func test() {
	s := "hello"
	c := []byte(s)
	c[0] = 'j'
	for i, j := range s {
		fmt.Printf("i = %d , j = %c\n", i, j)
	}

	for i, j := range string(c) {
		fmt.Printf("i = %d , j = %c\n", i, j)
	}

}
