package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("hello world;")
	a := "1"
	b := "111"

	res := addBinary(a, b)

	fmt.Println(res)
}

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	c := ""
	lenC := 0
	q := 0//进位
	sumTemp := 0
	for lenC < lenA || lenC < lenB {

		if lenC >= lenA {
			//证明A位数算完了
			res1, _ := strconv.Atoi(string(b[lenB - lenC - 1]))
			sumTemp = res1 + q
			if sumTemp >= 2 {
				sumTemp = 0
				q = 1
			} else {
				q = 0
			}
			c = strconv.Itoa(sumTemp) + c
			lenC += 1
		} else if lenC >= lenB {
		
			res1, _ := strconv.Atoi(string(a[lenA - lenC - 1]))
			sumTemp = res1 + q

			if sumTemp >= 2 {
				sumTemp = 0
				q = 1
			} else {
				q = 0
			}
			c = strconv.Itoa(sumTemp) + c

			lenC += 1
		} else {
			// fmt.Println(lenC)
			res1, _ :=  strconv.Atoi(string(a[lenA - lenC - 1]))
			res2, _ := strconv.Atoi(string(b[lenB - lenC - 1]))
			sumTemp = res1 + res2 + q
			q = 0
			if sumTemp >= 2 {
				//有进位
				sumTemp %= 2
				q = 1
			}
			c =  strconv.Itoa(sumTemp) + c
			lenC += 1
		}
	}
	if q == 1 {
		c =  "1" + c
	}

	return c
}