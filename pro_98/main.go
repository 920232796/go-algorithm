package main


import (
	"fmt"
)

//二叉搜索树，中序遍历为升序
func main() {

	root := &TreeNode{Val :5, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val:4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	
	res := isValidBST(root)
	fmt.Println(res)
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 
var last int64 = -100000

func isValidBST(root *TreeNode) bool {
	if root == nil {
        return true
    }
    if root.Left == nil && root.Right == nil {
        return true
    }
	res := dfs(root)
	return res
}
 
func dfs(node *TreeNode) bool {
	
	if node == nil {
		return true
	}

	if dfs(node.Left) {
		
		if int64(node.Val) < last {
			return false 
		}
		last = int64(node.Val)
		return dfs(node.Right)
	}
	return false
}