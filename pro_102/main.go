package main

import (
	"fmt"
)


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//二叉树的层次遍历，递归版本~
func main() {

	fmt.Println("hello world;")
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}}

	res := levelOrder(root)
	fmt.Println(res)

}

 
func levelOrder(root *TreeNode) [][]int {

	res := [][]int {}

	bfs(root, 0, &res)

	return res
}

func bfs(node1 *TreeNode, deepth int, res *[][]int) {

	if node1 == nil {
		return
	}

	if deepth >= len(*res) {
		*res = append(*res, []int{})
	}
	// fmt.Println(len(*res))
	(*res)[deepth] = append((*res)[deepth], node1.Val)

	bfs(node1.Left, deepth + 1, res)
	bfs(node1.Right, deepth + 1, res)
}