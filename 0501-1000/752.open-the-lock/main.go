// https://leetcode.cn/problems/open-the-lock/description/

package main

import (
	"fmt"
)

const Starter = "0000"

func openLock(deadends []string, target string) int {
	if Starter == target {
		return 0
	}

	visited := map[string]bool{}
	for _, v := range deadends {
		visited[v] = true
	}
	if visited[Starter] || visited[target] {
		return -1
	}
	return bfs(visited, target)
}

func bfs(visited map[string]bool, target string) int {
	step := map[string]int{Starter: 0}
	q := []string{Starter}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		next := make([]string, 8)
		var j int
		s := []byte(v)
		for i, b := range s {
			s[i] = ((b - 48 + 1) % 10) + 48
			next[j] = string(s)
			s[i] = ((b - 48 + 9) % 10) + 48
			next[j+1] = string(s)
			s[i] = b
			j += 2
		}

		for _, w := range next {
			if !visited[w] {
				visited[w] = true
				step[w] = step[v] + 1
				if w == target {
					return step[w]
				}
				q = append(q, w)
			}
		}
	}
	return -1
}

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}
