package main


import (
    "fmt"
)

func main() {

    nums := []int {3, 4, -1, 1}

    res := firstMissingPositive(nums)

    fmt.Println(res)

}

func firstMissingPositive(nums []int) int {

    res := make([]int, len(nums))

    for i := 0; i < len(nums); i ++ {

        if 0 < nums[i] && nums[i] <= len(nums) {

            res[nums[i] - 1] = nums[i]
        }
    }

    for i := 0; i < len(nums); i ++ {

        if (res[i] == 0) {
            return i + 1
        }
    }

    return len(nums) + 1

}
