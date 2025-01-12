package main

import "strconv"

func palindrome(x int) bool {
	xStr := strconv.Itoa(x)
	length := len(xStr)

	for i := 0; i < length/2; i++ {
		if xStr[i] != xStr[length-i-1] {
			return false
		}
	}

	return true
}
