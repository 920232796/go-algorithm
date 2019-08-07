package main

import (
    "fmt"
)

func main() {

    fmt.Println("hello world ")
    var li *ListNode
    li = &ListNode{Val: 2}
    li.Next = &ListNode{Val: 4}
    li.Next.Next = &ListNode{Val: 3}

    l2 := &ListNode{Val: 5}
    l2.Next = &ListNode{Val: 6}
    l2.Next.Next = &ListNode{Val: 4}

    result := addTwoNumbers(li, l2)

    res := []int {}

    for result != nil {
        res = append(res, result.Val)
        result = result.Next
    }

    for i := len(res) - 1; i >= 0; i -- {
        fmt.Print(res[i])
    }

}


type ListNode struct {
    Val int
    Next *ListNode
   }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode
    res = &ListNode{Val: 0}
    cur := res
    r := 0
    sum := 0

    for l1 != nil || l2 != nil {

        if l1 == nil {
            sum = l2.Val + r
            if sum >= 10 {
               r = 1
               sum = sum % 10
           } else {
               r = 0
           }
            l2 = l2.Next

        } else if l2 == nil {
            sum = l1.Val + r
            if sum >= 10 {
               sum = sum % 10
               r = 1
           } else {
               r = 0
           }
            l1 = l1.Next
        } else {
            sum = l1.Val + l2.Val + r
            if sum >= 10 {
                r = 1
                sum = sum % 10
            } else {
                r = 0
            }
            l1 = l1.Next
            l2 = l2.Next
        }

        cur.Next = &ListNode{Val: sum}
        cur = cur.Next


    }

    if r == 1 {
       //最后还有进位
       cur.Next = &ListNode{Val: 1}
   }

    return res.Next
}
