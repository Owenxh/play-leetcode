package main

import (
	"bytes"
)

// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
// 解释：
// P   A   H   N
// A P L S I I G
// Y   I   R

// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

func convert(s string, R int) string {
	if R >= len(s) || R == 1 {
		return s
	}
	sz := R + R - 2
	C := (len(s) + sz - 1) / sz * (R - 1)

	matrix := make([][]byte, R)
	for i := 0; i < R; i++ {
		matrix[i] = make([]byte, C)
	}

	x, y := 0, 0
	for i, ch := range s {
		matrix[x][y] = byte(ch)
		if i%sz < R-1 { // 向下移动
			x++
		} else { // 向右上移动
			y++
			x--
		}
	}

	ans := make([]byte, len(s))
	var i int
	for _, row := range matrix {
		for _, ch := range row {
			if ch != 0 {
				ans[i] = ch
				i++
			}
		}
	}
	return string(ans)
}

// 矩阵压缩

// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// 压缩后
// P I N
// A L S  I G
// Y A H R
// P I

func convert2(s string, R int) string {
	if R >= len(s) || R == 1 {
		return s
	}
	sz := R + R - 2
	matrix := make([][]byte, R)
	x := 0
	for i, ch := range s {
		matrix[x] = append(matrix[x], byte(ch))
		if i%sz < R-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(matrix, nil))
}

func main() {

}
