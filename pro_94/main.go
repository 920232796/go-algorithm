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

	fmt.Println("hello world;")
	
	root := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: nil}}

	res := inorderTraversal(root)
	fmt.Println(res)
}

 
func inorderTraversal(root *TreeNode) []int {
	res := []int {}
	
	dfs(root, &res)

	return res
	
}

func dfs(Node *TreeNode, res *[]int) {

	if Node == nil {
		return 
	}
	dfs(Node.Left, res)
	*res = append(*res, Node.Val)
	dfs(Node.Right, res)

}