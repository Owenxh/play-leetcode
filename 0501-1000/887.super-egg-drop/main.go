package main

import (
	"fmt"
	"sort"
)

func superEggDrop(K int, N int) int {
	// dp[i][j] 表示 i 层 j 个鸡蛋最少操作数
	f := make([][]int, N+1)
	for i := range f {
		f[i] = make([]int, K+1)
	}
	for i := 1; i <= N; i++ {
		for j := 1; j <= K; j++ {
			f[i][j] = i
		}
	}

	for i := 2; i <= N; i++ {
		for j := 2; j <= K; j++ {
			p := sort.Search(i+1, func(mid int) bool {
				// broken, unbroken := f[mid-1][j-1], f[i-mid][j]
				return f[mid-1][j-1] > f[i-mid][j]
			})
			p -= 1
			f[i][j] = max(f[p-1][j-1], f[i-p][j]) + 1
		}
	}
	for i := range f {
		for j, cnt := range f[i] {
			fmt.Printf("%2d 层楼楼 %2d 个鸡蛋最少操作数 %2d\n", i, j, cnt)
		}
	}
	return f[N][K]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	// fmt.Println(superEggDrop(1, 2))
	// fmt.Println(superEggDrop(2, 6))
	// fmt.Println(superEggDrop(3, 14))
	// fmt.Println(superEggDrop(10, 10000))

	superEggDrop(3, 14)
}
