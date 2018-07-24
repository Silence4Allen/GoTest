package util

import (
	"fmt"
	"GoTest/src/new"
)

func init() {
	fmt.Println("this is a test init")
}

const company string = "Tencent"
const bus = "公车"
const (
	cat  = "猫"
	dog
	fish string = "鱼"
)
const apple, banana string = "苹果", "香蕉"

func U1() {
	var (
		a int
		b int
		c int = 3
	)
	d := float32(c)
	fmt.Println(a + b + c)
	fmt.Println(d)
	fmt.Println("this is a test")
	new.U2()
}
