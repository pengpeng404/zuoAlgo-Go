package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 5}
	// 4
	fmt.Println(longestUnivaluePath(root))

}

// https://leetcode.com/problems/longest-univalue-path/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type info struct {
	length    int
	maxLength int
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return process(root).maxLength - 1
}

func process(x *TreeNode) *info {
	if x == nil {
		return &info{0, 0}
	}
	l := x.Left
	r := x.Right
	lInfo := process(l)
	rInfo := process(r)
	length := 1

	if l != nil && l.Val == x.Val {
		length = lInfo.length + 1
	}
	if r != nil && r.Val == x.Val {
		length = max(length, rInfo.length+1)
	}

	// 必须以 x 出发的最大路径信息有了 如上
	maxLength := max(max(lInfo.maxLength, rInfo.maxLength), length)
	if l != nil && r != nil && l.Val == x.Val && r.Val == x.Val {
		maxLength = max(maxLength, lInfo.length+rInfo.length+1)
	}
	return &info{length, maxLength}
}
