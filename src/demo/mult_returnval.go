package demo

import "fmt"

//编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
func MultTest() {
	m := 5
	n := 4
	show(getResult(m, n))
	show(getResult2(m, n))
}

func show(a, b, c int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func getResult(m, n int) (a, b, c int) {
	a = m + n
	b = m - n
	c = m * n
	return a, b, c
}
func getResult2(m, n int) (int, int, int) {
	return m + n, m - n, m * n
}
