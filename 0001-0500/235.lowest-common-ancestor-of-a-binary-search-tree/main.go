// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	l, r := p, q
	if l.Val > r.Val {
		l, r = r, l
	}
	return find(root, l, r)
}

func find(root, l, r *TreeNode) *TreeNode {
	if root.Val == l.Val {
		return l
	}
	if root.Val == r.Val {
		return r
	}
	if root.Val > r.Val {
		return find(root.Left, l, r)
	}
	if root.Val < l.Val {
		return find(root.Right, l, r)
	}
	return root
}
