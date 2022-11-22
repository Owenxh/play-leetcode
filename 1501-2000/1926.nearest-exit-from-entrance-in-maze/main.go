package main

import "fmt"

var dirs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func nearestExit(maze [][]byte, entrance []int) int {
	r, c := len(maze), len(maze[0])
	visited := make([][]bool, r)
	for i := 0; i < r; i++ {
		visited[i] = make([]bool, c)
	}
	var q [][3]int
	for i, row := range maze {
		for j, val := range row {
			if entrance[0] == i && entrance[1] == j {
				continue
			} else if (i == 0 || j == 0 || i == r-1 || j == c-1) && val == '.' {
				q = append(q, [3]int{i, j, 0})
				visited[i][j] = true
			} else if val == '+' {
				visited[i][j] = true
			}
		}
	}
	for len(q) > 0 {
		cur := q[0]
		x, y, dis := cur[0], cur[1], cur[2]
		q = q[1:]
		for i := 0; i < 4; i++ {
			nx, ny := x+dirs[i][0], y+dirs[i][1]
			if nx >= 0 && nx < r && ny >= 0 && ny < c && !visited[nx][ny] {
				if entrance[0] == nx && entrance[1] == ny {
					return dis + 1
				}
				visited[nx][ny] = true
				q = append(q, [3]int{nx, ny, dis + 1})
			}
		}
	}
	return -1
}

func main() {
	maze := [][]byte{{'+', '.', '+'}, {'.', '.', '.'}, {'+', '.', '+'}}
	fmt.Println(nearestExit(maze, []int{1, 2}))
}
