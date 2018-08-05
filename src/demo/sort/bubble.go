package sort

import "fmt"

func TestBubble() {
	data := IntSorter{19, 3, 65, 8, 35, 67, 43, 4, 99, 32}
	Sort(data)
	fmt.Println(data)
}

type Soter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(date Soter) {
	for pass := 1; pass < date.Len(); pass++ {
		for i := 0; i < date.Len()-pass; i++ {
			if date.Less(i, i+1) {
				date.Swap(i, i+1)
			}
		}
	}
}

type IntSorter []int

func (is IntSorter) Len() int {
	return len(is)
}
func (is IntSorter) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSorter) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}
