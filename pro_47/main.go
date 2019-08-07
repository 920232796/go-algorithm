
package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 2, 1, 1}

	result := permuteUnique(nums)

	fmt.Println(result)

}

func permuteUnique(nums []int) [][]int {
    res := [][]int{}

	permute_3(nums, 0, &res)

	fmt.Println(len(res))

	flag := 0

    for i := 0; i < len(res); i ++ {
        for j := i + 1; j < len(res); j ++ {

            for k := 0; k < len(res[i]); k ++ {
                if (res[i][k] == res[j][k]) {
                    flag += 1
                }
            }

            if (flag == len(res[i])) {
                //证明全部相同 要去掉
                res = append(res[:j], res[j + 1:]...)
			}

			flag = 0

        }
    }
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
