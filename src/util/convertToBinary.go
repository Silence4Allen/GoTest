package util

import "strconv"

func Run(n int) string {
	result := ""
	if n == 0 {
		result = "0"
	} else {

		for ; n > 0; n /= 2 {
			lsb := n % 2
			result = strconv.Itoa(lsb) + result
		}
	}
	return result
}
