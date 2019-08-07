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

	fmt.Println("hello world")

	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}}

	res := isSymmetric(root)
	fmt.Println(res)
}

func solve(node1 *TreeNode, node2 *TreeNode, flag *int) {
	

	if node1.Left == nil && node2.Right == nil {
		return 
	}

	if node1.Right == nil && node2.Left == nil {
		return 
	}

	if node1.Left == nil || node2.Right == nil {
		*flag += 1
		return 
	}

	if node1.Right == nil || node2.Left == nil {
		*flag += 1
		return 
	}

	if node1.Left.Val != node2.Right.Val {
		*flag += 1
		return
	} 
	if node1.Right.Val != node2.Left.Val {
		*flag += 1
		return 
	}
	solve(node1.Left, node2.Right, flag)
	solve(node1.Right, node2.Left, flag)
}



 func isSymmetric(root *TreeNode) bool {
	
	flag := 0

	solve(root.Left, root.Right, &flag) 

	if flag != 0 {
		return false
	} 
	return true
}