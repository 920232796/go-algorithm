package main


import (
	"fmt"

	)


func main() {
	fmt.Println("hello world;")

	nums := []int {1, 3, 5, 6}
	target := 5

	result := searchInsert(nums, target)

	fmt.Println(result)
}

func searchInsert(nums []int, target int) int {
	//二分查找？？
	left := 0
	right := len(nums) - 1
	var result int

	if (len(nums) == 1) {
		if (target > nums[0]) {
			return 1
		} else {
			return 0
		}
	}

	if (target > nums[right]) {
		result = right + 1
		return result
	}

	if (target == nums[right]) {
		return right
	}

	if (target <= nums[left]) {
		result = left
		return result
	}

	for left < right {

		if (right - left == 1) {
			result = right
			break
		} 
		middle := (right + left) / 2

		if (nums[middle] > target) {
			right = middle
		} else if (nums[middle] < target){
			left = middle
		} else {
			result = middle
			break
		}
	}
	return result

}
