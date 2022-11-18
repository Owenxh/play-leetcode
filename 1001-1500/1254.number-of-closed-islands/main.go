package main

func closedIsland(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	visited := make([][]bool, R)
	for i := 0; i < R; i++ {
		visited[i] = make([]bool, C)
	}

	pos := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var cnt int
	for i := 0; i < R; i++ {
		for j, v := range grid[i] {
			if !visited[i][j] && v != 1 {
				visited[i][j] = true
				if bfs(grid, R, C, i*C+j, visited, pos) {
					cnt++
				}
			}
		}
	}
	return cnt
}

func bfs(grid [][]int, R, C, start int, visited [][]bool, pos [][]int) bool {
	res := true
	var q []int
	q = append(q, start)
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		x, y := v/C, v%C
		if x == 0 || x == R-1 || y == 0 || y == C-1 {
			res = false
		}
		for i := 0; i < len(pos); i++ {
			nx, ny := x+pos[i][0], y+pos[i][1]
			if nx >= 0 && nx < R && ny >= 0 && ny < C && !visited[nx][ny] && grid[nx][ny] == 0 {
				visited[nx][ny] = true
				q = append(q, nx*C+ny)
			}
		}
	}
	return res
}
