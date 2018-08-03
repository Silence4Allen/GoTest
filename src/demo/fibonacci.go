package demo

import (
	"fmt"
)

const LIM = 2700

var fibs [LIM]uint64
func TestFibonacci() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
		fmt.Printf("f(%d) = %d\n", i, f())
	}
}

func fibo() func() int {
	x1, x2 := 0, 1
	sum := 0
	return func() int {
		sum = x1 + x2
		x1 = x2
		x2 = sum
		return sum
	}
}
func TestFibo2() {
	for i := 0; i < LIM; i++ {
		fibo2(i)
	}
}

func fibo2(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		fmt.Printf("已存在数组中，n = %d , res = %d\n", n, res)
		return
	}
	if n <= 1 {
		res = 1
		fmt.Printf("n = %d , res = %d\n", n, res)
	} else {
		res = fibo2(n-1) + fibo2(n-2)
		fmt.Printf("n = %d , res = %d\n", n, res)
	}
	fibs[n] = res
	fmt.Printf("把n = %d的值%d放入数组中\n", n, res)
	return
}

var arr3 [LIM]uint64

func TestFibo3() {
	fibo3()
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("i = %d , value = %d\n", i, arr3[i])
	}
}
func fibo3() {
	for i := 0; i < len(arr3); i++ {
		if i <= 1 {
			arr3[i] = 1
		} else {
			arr3[i] = arr3[i-1] + arr3[i-2]
		}
	}
}
