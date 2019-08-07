package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world;")

	lists := []*ListNode{}
	list1 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	list2 := &ListNode{Val: -10, Next: &ListNode{Val: -8, Next: &ListNode{Val: -5, Next: &ListNode{Val: -4}}}}
	// list5 := &ListNode{}
	// list6 := &ListNode{}
	list7 := &ListNode{Val: -3}	
	list3 := &ListNode{Val: -10, Next: &ListNode{Val: -9, Next: &ListNode{Val: -6, Next: &ListNode{Val: -4, Next: &ListNode{Val: -3, Next: &ListNode{Val: -2, Next: &ListNode{Val: -2, Next: &ListNode{Val: -1, Next: &ListNode{Val: 2}}}}}}}}}
	list4 := &ListNode{Val: -9, Next: &ListNode{Val: -9, Next: &ListNode{Val: -2, Next: &ListNode{Val: -1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}}
	list8 := &ListNode{Val: -9, Next: &ListNode{Val: -4, Next: &ListNode{Val: -3, Next: &ListNode{Val: -2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}}}}}}
	lists = append(lists, list1, list2, nil, nil, list7, list3, list4, list8)
	fmt.Println(lists)
	// list1 := &ListNode{} 
	// list2 := &ListNode{} 
	// lists = append(lists, nil, nil)
	res := mergeKLists1(lists)
	fmt.Println(res)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

type ListNode struct {
	Val int
    Next *ListNode
}

func mergeKLists1(lists []*ListNode) *ListNode {
	res := &ListNode{Val: 0}
	cur := res
	
	if len(lists) == 0 {
		return nil
	}

	//此时把k个链表的头元素构建成了最小堆！！ 下一步就是开始删掉头节点 插入新节点 下沉 再次构建最小堆
	// heapLen := len(lists)
	for len(lists) != 0 {
		lists = consHeap1(lists)
		//heapLen != 0 证明堆还没空 还要继续循环
		if len(lists) == 0 {
			return nil
		}
		cur.Next = lists[0]//拿出头上最小的 
		
		cur = cur.Next
		if lists[0].Next != nil {
			lists = append(lists, lists[0].Next)
		
		}
	
		lists = lists[1: ]
	
	}
    return res.Next
}

func consHeap1(lists []*ListNode) []*ListNode {

	
	for i := 0; i < len(lists); i ++ {
		fmt.Printf("lists[i]: %v\n", lists[i])
        if lists[i] == nil {
			fmt.Println("delite!!")
			lists = append(lists[:i], lists[i + 1:]...)
			i -= 1
        }
	}
	
	

	fmt.Printf("lists: %v\n", lists)
	
	
	var j int = (len(lists) - 1 ) / 2
	for ; j >= 0; j -- {

		if 2 * j + 2 <= len(lists) - 1 && lists[j].Val > lists[2 * j + 2].Val {
			//证明最后一个根节点有右孩子
			lists[j], lists[2 * j + 2] = lists[2 * j + 2], lists[j]
		}
		if 2 * j + 1 <= len(lists) - 1 && lists[j].Val > lists[2 * j + 1].Val {
			//zh证明有左孩子
			lists[j], lists[2 * j + 1] = lists[2 * j + 1], lists[j]
		}
	}
	
	fmt.Println(lists)
	return lists 

}


