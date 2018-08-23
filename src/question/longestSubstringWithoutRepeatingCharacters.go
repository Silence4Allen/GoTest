package question

import "fmt"

/**
给定一个字符串，找出不含有重复字符的最长子串的长度。
 */
func LongestSubstringWithoutRepeatingCharacters(s string) (int, string) {
	m := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := m[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		m[ch] = i
	}
	return maxLength, string(s[start : start+maxLength])
}

func PrintLongestSubstringWithoutRepeatingCharactersResult(str string) {
	maxLength, subStr := LongestSubstringWithoutRepeatingCharacters(str)
	fmt.Printf("maxLength = %d , subStr = %s\n", maxLength, subStr)
}
