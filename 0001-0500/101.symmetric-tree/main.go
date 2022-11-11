// https://leetcode.cn/problems/symmetric-tree/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(left, right *TreeNode) bool {
	if !isSame(left, right) {
		return false
	}
	if left == nil {
		return true
	}

	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

func isSame(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	return p.Val == q.Val
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Print(isSymmetric(root))
}
