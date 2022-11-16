package main

import "fmt"

type Solution struct {
	Col   []bool
	Diag1 []bool
	Diag2 []bool
	Ans   [][]string
}

func solveNQueens(n int) [][]string {
	s := Solution{
		Col:   make([]bool, n),
		Diag1: make([]bool, 2*n-1),
		Diag2: make([]bool, 2*n-1),
	}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	backtrack(&s, n, 0, queens)
	return s.Ans
}

// 解决 N 皇后问题，放置第 i 个皇后
func backtrack(s *Solution, n int, i int, queens []int) {
	// 递归到底，找到 N 个皇后的一个解
	if i == n {
		//fmt.Println("Find out N queens result:", queens)
		s.Ans = append(s.Ans, generateBoard(queens, n))
		return
	}

	//fmt.Printf("Try to put %v queen\n", i)

	for j := 0; j < n; j++ {
		if s.Col[j] || s.Diag1[i+j] || s.Diag2[i-j+n-1] {
			continue
		}
		s.Col[j] = true
		s.Diag1[i+j] = true
		s.Diag2[i-j+n-1] = true
		queens[i] = j

		//fmt.Printf("Put [%v] queen at column [%v],  diagonal1 [%v], diagonal2 [%v] \n", i, j, i+j, i-j+n-1)

		backtrack(s, n, i+1, queens)
		queens[i] = -1
		s.Col[j] = false
		s.Diag1[i+j] = false
		s.Diag2[i-j+n-1] = false
	}
}

func generateBoard(queens []int, n int) []string {
	board := make([]string, n)
	for i, queen := range queens {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if j == queen {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		board[i] = string(row)
	}
	return board
}

func main() {
	n := 8
	//boards := solveNQueens(4)
	boards := solveNQueens(n)
	fmt.Printf("total %v solution for %v queens\n", len(boards), n)
	for _, b := range boards {
		for _, row := range b {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
