package main


import (
  "fmt"
)

func main() {

  fmt.Println("hello world;")

  nums := []int {3, 2, 2, 3}

  val := 3

  res := removeElement(nums, val)

  fmt.Println(res)
}

func removeElement(nums []int, val int) int {

    for i := 0; i < len(nums); i ++ {

    if (nums[i] == val) {
      nums = append(nums[:i], nums[i + 1:]...)
      i = i - 1
    }
  }

  return len(nums)
}
