package main

import "fmt"

func main() {
	candidates := []int {2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {

	//回溯法
	var res [][]int
	cur := []int{}
	backTrace(candidates, target, cur, &res)
	return res

}

func backTrace(candidates []int, target int, cur []int, res *[][]int) {

	//递归停止条件
	if (target == 0) {
		temp := []int {}
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	if (target < 0) {
		return
	}

	for i := 0; i < len(candidates); i ++ {
		cur = append(cur, candidates[i])
		backTrace(candidates, target - candidates[i], cur, res)
		cur = cur[: len(cur) - 1]
	}

}