package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("hello wolrd;")
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 0}}
	res := sumNumbers(root)
	fmt.Println(res)
}

func sumNumbers(root *TreeNode) int {
	res := 0
	cur := ""
	dfs(root, &res, cur)
	return res 
}

func dfs(node *TreeNode, res *int, cur string) {
	cur = cur + strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		temp, _ := strconv.Atoi(cur)
		*res = *res + temp
		return 
	}
	dfs(node.Left, res, cur)
	dfs(node.Right, res, cur)

}