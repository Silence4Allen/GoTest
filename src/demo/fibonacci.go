package demo

import "fmt"

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

func fibo2(n int) (res int) {
	if n <= 1 {
		res = 1
	}
	return fibo2(n-1) + fibo2(n-2)
}
