package demo

import (
	"math"
	"errors"
	"fmt"
)

//编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值。
func ErrorTest() {
	fmt.Print("First example with -1: ")
	ret1, err1 := mysqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Print("Second example with 5: ")
	//you could also write it like this
	if ret2, err2 := mysqrt2(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
}

func mysqrt(f float64) (result float64, err error) {
	if f < 0 {
		result = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		result = math.Sqrt(f)
	}
	return
}

func mysqrt2(f float64) (float64, error) {
	if f < 0 {
		return math.NaN(), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(f), nil
}
