package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	
	fmt.Println("hello wolrd;")

	root := &TreeNode{Val: -10, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: -10, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: -7}}}

	res := maxPathSum(root)

	fmt.Println(res)
}

func maxPathSum(root *TreeNode) int {
	
	Max := -1000
	solve(root, &Max)

	return Max
}

func solve(node *TreeNode, Max *int) int {
	
	if node == nil {
		return 0
	}

	left := solve(node.Left, Max)
	right := solve(node.Right, Max)

	// *Max = max(max(*Max, max(left, 0) + max(right, 0) + node.Val), max(left, right))
	*Max = max(*Max, max(left, 0) + max(right, 0) + node.Val)
	return max(0, max(left, right)) + node.Val 
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}