package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello world;")

	digits := []int {9, 9, 9, 9}
	res := plusOne(digits)

	fmt.Println(res)

}

func plusOne(digits []int) []int {

    for i := len(digits) - 1; i >= 0; i -- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			//证明有进位
			digits[i] = 0
		} else {
			break
		}
	}
	
	if digits[0] == 0 {
		// 证明第一位也进位！
		digits[0] = 1
		digits = append(digits, 0)
	}
	
	
	// digits = append(digits, 1)
	return digits

}