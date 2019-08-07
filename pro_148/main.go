package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	res := sortList(head)
	
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

func sortList(head *ListNode) *ListNode {

	if head.Next == nil {
		//证明只有一个元素了!!
		return head 
	}
	middle := getMiddle(head)

	nextMiddle := middle.Next
	middle.Next = nil

	
	resLeft := sortList(head)
	resRight := sortList(nextMiddle)

	return merge(resLeft, resRight)

}

func getMiddle(head *ListNode) *ListNode {
	if head == nil {
		return head 
	}
	slow := head
	fast := head.Next

	if fast != nil {
		fast = fast.Next 
		if fast != nil {
			slow = slow.Next 
			fast = fast.Next 
		}
	}
	return slow
}

func merge(left *ListNode, right *ListNode) *ListNode {

	if left == nil {
		return right 
	} else if right == nil {
		return left 
	} 
	if left.Val <= right.Val {
		left.Next = merge(left.Next, right)
		return left
	} 
	if left.Val > right.Val {
		right.Next = merge(left, right.Next)
		return right 
	}
	return nil 
}


