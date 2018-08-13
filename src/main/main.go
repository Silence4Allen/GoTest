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
	fmt.Printf("\nIt takes %v to run this program.\n", a)

}

func test() {
	reboot := aa{
		"reboot",
		21,
	}

	people := aa{
		"allen",
		33,
	}

	fmt.Println(reboot)
	fmt.Println(people)
	changeAge(&reboot)
	fmt.Println(reboot)
}
func changeAge(a *aa) {
	a.age = 100
}

type aa struct {
	name string
	age  int
}
