package main

type graph struct {
	G [][]int
	R int
	C int
}

func maxAreaOfIsland(grid [][]int) int {
	R, C := len(grid), len(grid[0])

	g := &graph{
		G: grid,
		R: R,
		C: C,
	}

	var ret int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 1 {
				ret = max(ret, dfs(g, i, j))
			}
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func dfs(g *graph, i, j int) int {
	if i < 0 || i >= g.R || j < 0 || j >= g.C {
		return 0
	}

	if g.G[i][j] == 0 {
		return 0
	}

	g.G[i][j] = 0
	res := 1

	res += dfs(g, i-1, j)
	res += dfs(g, i, j+1)
	res += dfs(g, i+1, j)
	res += dfs(g, i, j-1)

	return res
}
