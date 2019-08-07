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
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}

	flatten(root)

	dfs(root)
	// fmt.Println(root)

	// fmt.Println(root.Left)
	// fmt.Println(root.Left)



}

func flatten(root *TreeNode)  {
	fmt.Println("hello word")
	solve(root)

	
}

func solve(node *TreeNode) {
	
	if node == nil {
		return 
	}

	solve(node.Left)
	solve(node.Right)
	
	if node.Left != nil && node.Right != nil {
		cur := node.Left 
		
		for cur.Right != nil {
			cur = cur.Right 
		}
		cur.Right = node.Right
		node.Right = node.Left
		node.Left = nil 
	} else if node.Left != nil && node.Right == nil {
		node.Right = node.Left
		node.Left = nil 
	}


}

func dfs(node *TreeNode) {
	if node == nil {
		fmt.Println("nil")
		return 
	}

	fmt.Println(node.Val)
	dfs(node.Left)
	dfs(node.Right)

}