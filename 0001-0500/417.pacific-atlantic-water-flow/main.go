package main

import (
	"fmt"
)

var dirs = []struct{ x, y int }{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(int, int, [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		for _, d := range dirs {
			if nx, ny := x+d.x, y+d.y; nx >= 0 && nx < m && ny >= 0 && ny < n && heights[nx][ny] >= heights[x][y] {
				dfs(nx, ny, ocean)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for j := 1; j < n; j++ {
		dfs(0, j, pacific)
	}
	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}
	for j := 0; j < n-1; j++ {
		dfs(m-1, j, atlantic)
	}
	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}

func main() {
	heights := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	fmt.Println(pacificAtlantic(heights))
}
