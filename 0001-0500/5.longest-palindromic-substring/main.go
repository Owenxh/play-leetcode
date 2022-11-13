package main

import "fmt"

func longestPalindrome(s string) string {
	arr := []byte(s)
	for length := len(arr); length >= 1; length-- {
		for l := 0; l+length <= len(arr); l++ {
			r := l + length - 1
			if isLP(arr, l, r) {
				return string(arr[l : r+1])
			}
		}
	}
	return ""
}

func isLP(s []byte, l, r int) bool {
	for i, j := l, r; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
