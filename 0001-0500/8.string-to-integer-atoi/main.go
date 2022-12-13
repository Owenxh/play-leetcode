package main

import (
	"fmt"
	"math"
	"strconv"
)

func parse(s string) []byte {
	var arr []byte
	for _, b := range []byte(s) {
		switch {
		case b >= '0' && b <= '9':
			if len(arr) == 0 {
				arr = append(arr, '+')
			}
			arr = append(arr, b)
		case (b == '+' || b == '-') && len(arr) == 0:
			arr = append(arr, b)
		case b == ' ' && len(arr) == 0:
			continue
		default:
			return arr
		}
	}
	return arr
}

func myAtoi(s string) int {
	arr := parse(s)
	if len(arr) == 0 {
		return 0
	}
	neg := arr[0] == '-'
	arr = arr[1:]
	for len(arr) > 0 && arr[0] == '0' {
		arr = arr[1:]
	}
	str := string(arr)
	if len(str) > 10 {
		if neg {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	} else if len(str) == 10 {
		if neg {
			if str >= strconv.Itoa(math.MinInt32)[1:] {
				return math.MinInt32
			}
		} else {
			if str >= strconv.Itoa(math.MaxInt32) {
				return math.MaxInt32
			}
		}
	}
	var ans int
	for i, b := range arr {
		num := int(b - '0')
		ans += num * int(math.Pow(10, float64(len(arr)-1-i)))
	}
	if neg {
		return -ans
	} else {
		return ans
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	strs := []string{
		"   -42",
		"-2147483647",
		"4193 with words",
		"words and 987",
		"-91283472332",
		"+1",
		"-2147483647",
		"20000000000",
	}
	for _, s := range strs {
		fmt.Printf("parse %25v to %d\n", s, myAtoi(s))
	}
}
