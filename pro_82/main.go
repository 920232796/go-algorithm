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

	head := &ListNode{Val:10, Next: &ListNode{Val: 5, Next: &ListNode{Val: 20}}}

	cur := head

	cur.Next = nil

	fmt.Println(head)
}

