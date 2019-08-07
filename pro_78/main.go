package main


import (
	"fmt"
)

func main() {
	fmt.Println("hello world;")
	res := subsets([]int {1, 2, 3})
	fmt.Println(res)
}

func subsets(nums []int) [][]int {

	res := [][]int {}
	cur := []int {}
	
	solve(&res, cur, nums, 0)
	return res
    
}

func solve(res *[][]int, cur []int, nums []int, begin int) {

	if begin > len(nums) {
		return 
	}

	temp := make([]int, len(cur))
	copy(temp, cur)
	*res = append(*res, temp)

	for i := begin; i < len(nums); i ++ {
		cur = append(cur, nums[i])
		// fmt.Println(cur)
		solve(res, cur, nums, i + 1)
		cur = cur[: len(cur) - 1]
	}

}