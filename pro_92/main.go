package main

import (
	"fmt"
)


type ListNode struct {
	Val int
	Next *ListNode
}
 
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n {
		return head 
	}
	dummy := &ListNode{Val: -1, Next: head} //定义这个就是为了m = 1的时候，如果直接用head m = 1 根本找不到前驱节点！
	pre := dummy
	for i := 0; i < m - 1; i ++ {
		pre = pre.Next//找到 需要开始反转的那个node的前驱节点
		
	}
	cur := pre.Next
	for j := m - 1; j < n - 1; j ++ {
		t := cur.Next 
		cur.Next = t.Next 
		t.Next = pre.Next
		pre.Next = t
	}

	return dummy.Next 
    
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	res := reverseBetween(head, 2, 4)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	// fmt.Println("hello world")
}