// https://leetcode.cn/problems/shortest-path-in-binary-matrix/description/

package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	R, C := len(grid), len(grid[0])
	pos := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	return bfs(pos, grid, R*C-1, R, C)
}

func bfs(pos [][]int, G [][]int, end, R, C int) int {
	G[0][0] = 1
	queue := []int{0, 1}
	for len(queue) > 0 {
		v, step := queue[0], queue[1]
		if v == end {
			return step
		}
		queue = queue[2:]
		i, j := v/C, v%C

		for _, p := range pos {
			m, n := i+p[0], j+p[1]
			if m >= 0 && m < R && n >= 0 && n < C && G[m][n] == 0 {
				G[m][n] = 1
				w := m*C + n
				queue = append(queue, w, step+1)
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
}
