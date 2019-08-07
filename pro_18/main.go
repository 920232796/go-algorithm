package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("hello world;")
	nums := []int { 0, -1, -3, 5, -5}
	// nums := []int {-1, 0, 1, 2, -1, -4}
	// nums := []int {1, 0, -1, 0, -2, 2}
	// nums := []int {0, 0, 0, 0}
	// nums := []int {-4, 0, -4, 2, 2, 2 , -2, -2}
	target := 1
	res := fourSum(nums, target)

	fmt.Println(res)
}

func fourSum(nums []int, target int) [][]int {
	res := [][]int {}
 
 //判断长度
 if len(nums) < 4 {
	 return res
 }

 if len(nums) == 4 {
	 if nums[0] + nums[1] + nums[2] + nums[3] == target {
		 res = append(res, nums)
		 return res
	 }
 }
 
 sort.Ints(nums)
 l := 0
 r := 0
 

 for i := 0; i <= len(nums) - 4; i ++ {

	 if  i > 0 && nums[i] == nums[i - 1] {
		 continue
	 }
	 
	 for j := i + 1; j <= len(nums) - 3; j ++ {

	
		 if j > i + 1 && nums[j] == nums[j - 1] {
			 continue
		 }

		 l = j + 1
		 r = len(nums) - 1
		 
		 for l < r {
			
			fmt.Println(i, j, l, r)
			if nums[i] + nums[j] + nums[l] + nums[r] < target {

				 l += 1
				 
				 for l < r && nums[l] == nums[l - 1]   {
					 l += 1
			 
				 }

				 

			 } else if nums[i] + nums[j] + nums[l] + nums[r] > target {

				 r -= 1
				 for l < r && nums[r] == nums[r + 1] {
					 r -= 1
				 
				 }
				 
			 } else {

				 res = append(res, []int {nums[i], nums[j], nums[l], nums[r]})
			 
				 r -= 1
				 l += 1
				 
				 for l < r && nums[r] == nums[r + 1]  {
					 r -= 1
					 
				 }
				 
				 for l < r && nums[l] == nums[l - 1] {
					 l += 1
					 
				 }
				 
				 
				 
			 }
		 }
		 

	 }


	 
 }
 return res
 
}