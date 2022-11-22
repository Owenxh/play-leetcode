package main

var dirs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func shortestBridge(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	sx, sy := calFirstLandPoint(grid)
	grid[sx][sy] = 2

	var q [][3]int
	land := [][2]int{{sx, sy}}
	for len(land) > 0 {
		x, y := land[0][0], land[0][1]
		q = append(q, [3]int{x, y, 0})
		land = land[1:]
		for i := 0; i < 4; i++ {
			nx, ny := x+dirs[i][0], y+dirs[i][1]
			if nx >= 0 && nx < r && ny >= 0 && ny < c && grid[nx][ny] == 1 {
				grid[nx][ny] = 2
				land = append(land, [2]int{nx, ny})
			}
		}
	}
	ans := -1
	for len(q) > 0 {
		cur := q[0]
		x, y, dis := cur[0], cur[1], cur[2]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := x+dirs[i][0], y+dirs[i][1]
			if nx >= 0 && nx < r && ny >= 0 && ny < c && grid[nx][ny] != 2 {
				if grid[nx][ny] == 1 {
					return dis
				}
				grid[nx][ny] = 2
				q = append(q, [3]int{nx, ny, dis + 1})
			}
		}
	}
	return ans
}

func calFirstLandPoint(grid [][]int) (int, int) {
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				return i, j
			}
		}
	}
	return -1, -1
}
