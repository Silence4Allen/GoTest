package main

import "fmt"

func main() {
	a := 5
	p := &a
	*p = 7

	fmt.Printf("p代表的是a的地址：%x\n", p)
	fmt.Printf("&a代表的是a的地址：%x\n", &a)
	fmt.Printf("*p代表的是a的值：%d\n", *p)
	fmt.Printf("a代表的是a的值：%d\n", a)
	fmt.Printf("&p代表的是p的地址：%x\n", &p)

	s := []int{7, 9, 3, 5, 1}
	for index, content := range s {
		fmt.Printf("index is %d 的为：", index)
		fmt.Println(content)
	}
}
