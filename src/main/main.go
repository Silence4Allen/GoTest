package main

import (
	"fmt"
	"time"
	"GoTest/src/question"
)

func main() {
	start := time.Now()
	test()
	end := time.Now()
	a := end.Sub(start)
	fmt.Printf("\nIt takes %v to run this program.\n", a)

}

func test() {
	question.PrintLongestSubstringWithoutRepeatingCharactersResult("hdlajljfdal")
	question.PrintLongestSubstringWithoutRepeatingCharactersResult("d")
	question.PrintLongestSubstringWithoutRepeatingCharactersResult("")
	question.PrintLongestSubstringWithoutRepeatingCharactersResult("abcdefghijklmn")
	question.PrintLongestSubstringWithoutRepeatingCharactersResult("Process finished with exit cod")
}


