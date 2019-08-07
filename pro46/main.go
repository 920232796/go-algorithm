package main


import (
	"fmt"
)

func main() {

	nums := []int{1, 1, 2}

	result := permute(nums)

	fmt.Println(result)

	fmt.Println("hello world;")
}



func permute(nums []int) [][]int {

	res := [][]int{}

	permute_3(nums, 0, &res)

	return res

}

func permute_3(nums []int, k int, res *[][]int) {

	if (k == len(nums)) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}

	for i := k; i < len(nums); i ++ {
		//k就是被固定的那个元素 然后 交换 之后继续对后面剩下的递归
		temp := nums[i]
		nums[i] = nums[k]
		nums[k] = temp

		//对后面剩下的递归

		permute_3(nums, k + 1, res)

		//换回来还要
		temp = nums[i]
		nums[i] = nums[k]
		nums[k] = temp
	}
}
