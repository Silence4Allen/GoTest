package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	Handler()
	fmt.Println(time.Since(start))
}

func Handler() {
	time.Sleep(time.Second * 10)
	fmt.Println("hello")
}
