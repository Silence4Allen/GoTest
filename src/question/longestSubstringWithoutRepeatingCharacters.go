package question

import "fmt"

func LongestSubstringWithoutRepeatingCharacters(s string) (int, string) {
	m := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
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
