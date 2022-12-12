package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func decodeString(s string) string {
	src := []rune(s)
	substr, _ := solve(src, 0, 1)
	return substr
}

func calcString(src []rune, l int) (string, int) {
	i := l
	for i < len(src) && unicode.IsLetter(src[i]) {
		i++
	}
	return string(src[l:i]), i
}

func calcNumber(src []rune, l int) (int, int) {
	i := l
	for i < len(src) && unicode.IsDigit(src[i]) {
		i++
	}
	if src[i]-'[' != 0 {
		panic(fmt.Sprintf("index of %d must be char '['", i))
	}
	cnt, err := strconv.Atoi(string(src[l:i]))
	if err != nil {
		panic(err)
	}
	return cnt, i + 1
}

func solve(src []rune, l int, repeat int) (string, int) {
	var res string
	i := l
	for i < len(src) && src[i] != ']' {
		if unicode.IsLetter(src[i]) {
			substr, p := calcString(src, i)
			res += substr
			i = p
		} else if unicode.IsDigit(src[i]) {
			cnt, p := calcNumber(src, i)
			var substr string
			substr, i = solve(src, p, cnt)
			res += substr
		}
	}
	var ans string
	for k := 0; k < repeat; k++ {
		ans += res
	}
	return ans, i + 1
}

func main() {
	str := "3[z]2[2[y]pq4[2[jk]e1[f]]]ef"
	// "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"
	fmt.Println(decodeString(str))
}
