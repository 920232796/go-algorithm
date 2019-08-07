package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world;")

	lists := []int {2, 1, 4, 3, 6, 4, 10, 9, 2}
	fmt.Println(lists)
	res := mergeKLists(lists)
	fmt.Println(res)
}

func mergeKLists(lists []int) []int  {

	headLen := len(lists)

	res := []int {}

	for headLen != 0 {

		lists = consHeap(lists)
		res = append(res, lists[0])
		fmt.Println(res)
		lists = lists[1:]
		headLen -= 1
	}

    return res
}

func consHeap(lists []int) []int {

	var j int = (len(lists) - 1 ) / 2
	for ; j >= 0; j -- {

		if 2 * j + 2 <= len(lists) - 1 && lists[j] > lists[2 * j + 2] {
			//证明最后一个根节点有右孩子
			lists[j], lists[2 * j + 2] = lists[2 * j + 2], lists[j]
		}
		if 2 * j + 1 <= len(lists) - 1 && lists[j] > lists[2 * j + 1] {
			//zh证明有左孩子
			lists[j], lists[2 * j + 1] = lists[2 * j + 1], lists[j]
		}
	}
	
	return lists 

}
