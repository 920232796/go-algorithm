package main

import (
	"fmt"
	"strconv"
)

//先求所有的排列 再求第几个排列！

func main() {

	res := getPermutation(3, 5)

	fmt.Println(res)

	// fmt.Println("hello world;")
}

func getPermutation(n int, k int) string {
	pList := [][]int {}
	cur := []int {}

	for i := 1; i <= n; i ++ {
		cur = append(cur, i)
	}

	// fmt.Println(cur)
	// fmt.Println(k)
	permutation(&pList, cur, 0)

	fmt.Println(pList)

	res := ""
	for _, v := range(pList[k - 1]) {
		res += strconv.Itoa(v)
	}


	return res

}

func permutation(pList *[][]int, cur []int, begin int) {

	if len(cur) == begin {
		//证明够了 就添加一个结果
		temp := make([]int, len(cur))
		copy(temp, cur)
		*pList = append(*pList, temp)
		return
	}

	for i := begin; i < len(cur); i ++ {
		//固定一个元素  然后不断的让每个元素跟它交换 递归下去 一直到固定最后一个元素了 
		cur[i], cur[begin] = cur[begin], cur[i]
		permutation(pList, cur, begin + 1)
		cur[i], cur[begin] = cur[begin], cur[i]
	}

}