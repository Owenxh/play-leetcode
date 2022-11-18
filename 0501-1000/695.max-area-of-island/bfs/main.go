package dfs

func maxAreaOfIsland(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	pos := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	var ans int
	for i, vertexs := range grid {
		for j, v := range vertexs {
			if v == 1 {
				ans = max(ans, bfs(grid, R, C, i, j, pos))
			}
		}
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func bfs(grid [][]int, R, C, i, j int, pos [][]int) (res int) {
	grid[i][j] = 0
	q := []int{i*C + j}
	for len(q) > 0 {
		w := q[0]
		x, y := w/C, w%C
		res += 1
		q = q[1:]
		for i := 0; i < len(pos); i++ {
			nx, ny := x+pos[i][0], y+pos[i][1]
			if nx >= 0 && nx < R && ny >= 0 && ny < C && grid[nx][ny] == 1 {
				grid[nx][ny] = 0
				q = append(q, nx*C+ny)
			}
		}
	}
	return
}
