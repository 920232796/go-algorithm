package main

import (
    "fmt"
    "sort"
    "math"
)

func main() {

    fmt.Println("hello world;")

    nums := []int {0, 2, 1, -3}
    target := 1
    res := threeSumClosest(nums, target)

    // i := -1000
    // fmt.Println(math.Abs(float64(i)))

    fmt.Println(res)

}


func threeSumClosest(nums []int, target int) int {

    sort.Ints(nums)
    minD := 10000
    res := 0
    absValue := 0
    value := 0
    var l int
    var r int

    for i := 0; i < len(nums) - 2; i ++ {

        l = i + 1
        // fmt.Println(l)
        r = len(nums) - 1


        for l < r {
            value = target - (nums[i] + nums[l] + nums[r])
            absValue = int(math.Abs(float64(value)))
            // fmt.Printf("value : %d", value)
            // fmt.Printf("i: %d", i)
            // fmt.Printf("l r : %d, %d\n", l, r)

            if absValue >= minD {
                if value > 0 {
                    //证明有点大 右边大 所以左侧左移
                    l += 1
                } else if value < 0 {
                    r -= 1
                } else {
                    //证明为0 最接近了 直接返回就行了
                    return nums[i] + nums[l] + nums[r]
                }
            } else if absValue < minD {
                //证明找到更接近的了
                minD = absValue
                // fmt.Printf("minD: %d", minD)
                res = nums[i] + nums[l] + nums[r]
                // fmt.Printf("res: %d\n", res)
                if value < 0 {
                    //这时候是正的，因此r右移没意义了 只能更大 可以l左移
                    r -= 1
                } else if value > 0 {
                    l += 1
                } else {
                    return nums[i] + nums[l] + nums[r]
                }
            }
        }

    }

    return res
    // return 0
}
