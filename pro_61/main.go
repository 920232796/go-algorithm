package main

import (
	"fmt"
)


type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

	fmt.Println("hello world;")

	head := &ListNode{Val:0}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	// head.Next.Next.Next = &ListNode{Val: 4}
	
	// for head != nil {
	// 	fmt.Println(head)
	// 	head = head.Next
	// }

	res := rotateRight(head, 4)

	// fmt.Println(res)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
	
}

func rotateRight(head *ListNode, k int) *ListNode {
    
    if k == 0 || head == nil {
        return head
    }
    cur := head

	L := 1

	for cur.Next != nil {
		L += 1
		cur = cur.Next
	}
	//搞成了循环链表了
	cur.Next = head

	cur = cur.Next

	//计算循环次数 把目标位置的next整成nil就行了
	loop := L - k % L

	for j := 0; j < loop - 1; j ++ {
		cur = cur.Next
	}
	end := cur.Next
	cur.Next = nil

	return end
}