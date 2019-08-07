package main

import (
	"fmt"
)

//对链表进行插入排序


 
type ListNode struct {
	Val int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	
	dummy := &ListNode{Val: -1}
	dummy.Next = head 

	for head != nil  && head.Next != nil {
		if head.Next.Val > head.Val {
			//不需要动
			// fmt.Println("不需要动")
			head = head.Next 
			continue
		} else {
			//for 循环 从头往后找，找到第一个比当前node大的那个node 插入到这个node前面
			pre := dummy
			for pre.Next != nil {
				if head.Next.Val < pre.Next.Val {
					//证明找到啦 head.Next这个node要换到 pre 这个node之后 pre.Next 这个node 之前
					cur := head.Next 
					head.Next = cur.Next 
					cur.Next = pre.Next 				
					pre.Next = cur
					break
				}
				pre = pre.Next 
			}
		}
	} 

	return dummy.Next
	
}

func main() {

	head := &ListNode {Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}

	// for head != nil {
	// 	fmt.Println(head.Val)
	// 	head = head.Next 
	// }

	res := insertionSortList(head)
	
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next 
	}
}