package main

import (
    "fmt"
    "myFunc/aboutNum"
)


func main() {

    height := []int {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

    res := trap(height)

    fmt.Println(res)
}

func trap(height []int) int {
    if len(height) < 3 {
        //说明存不了水
        return 0
    }

    //面积
    res := 0

    for i := 1; i < len(height) - 1; i ++ {
        left := i - 1
        right := i + 1
        temp := aboutNum.Min(height[left], height[right]) - height[i]
        //判断 如果两边最小的那一个 减去中间那个 结果都大于0 证明可以存水!
        if temp > 0 {
            //证明可以存水
            res += temp
            //继续向两边循环 直到left - 1 < left  right + 1 > right 停止
            for height[left - 1] > height[left] {
                left -= 1
                res +=
            }
            for height[right + 1] > height[right] {
                right += 1
            }
        }
    }

}
