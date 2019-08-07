package main


import (
	"fmt"
)

func main() {
	fmt.Println("hello world;")
	res := combine(4, 2)
	fmt.Println(res)
}

func combine(n int, k int) [][]int {

	res := [][]int {}
	list := make([]int, n)
	cur := []int {}
	for i := 0; i < n; i ++ {
		list[i] = i + 1
	}

	solve(&res, list, cur, 0, k)
	return res
    
}

func solve(res *[][]int, list []int, cur []int, begin int, k int) {

	if len(cur) == k {
		//证明找到了一个结果
		temp := make([]int, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return 
	}

	for i := begin; i < len(list); i ++ {
		cur = append(cur, list[i])
		// fmt.Println(cur)
		solve(res, list, cur, i + 1, k)
		cur = cur[: len(cur) - 1]
	}

}