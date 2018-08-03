package demo

import "fmt"

func TestMap() {
	getSlice()
}

func mapDemo1() {
	mf := map[int]func() int{
		1: demo1Sp,
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
	for i, j := range mf {
		fmt.Printf("i = %d , j = %d\n", i, j())
	}
}
func demo1Sp() int {
	return 33
}

//注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的
func mapDemo2() {
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}
}

func getSlice() {
	slices := make([]map[int]int, 5)
	for _, j := range slices {
		j = make(map[int]int, 2)
		j[0] = 1
	}
	fmt.Printf("Version A: Value of items: %v\n", slices)
	fmt.Printf("------------------\n")
	for i := range slices {
		slices[i] = make(map[int]int, 2)
		slices[i][0] = 1
	}
	fmt.Printf("Version A: Value of items: %v\n", slices)
}
