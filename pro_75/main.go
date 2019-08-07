package main


import (
	"fmt"
)

func main() {

	fmt.Println("hello world;")
	nums := []int {2, 0, 2, 1, 1, 0}
	sortColors(nums)

}

func sortColors(nums []int)  {

		for i := 0; i < len(nums); i ++ {
			if nums[i] == 0 {
				nums = append(nums[: i], nums[i + 1: ]...)
				nums = append([]int {0}, nums... )
			
			} else if nums[i] == 2 {
				
				
				nums = append(nums[: i], nums[i + 1: ]...)
				nums = append(nums, 2)
				for nums[i] == 2 && i < len(nums) - 1 {
					i = i + 1	
				}
				
			}
			
		}
		fmt.Println(nums)
		
}