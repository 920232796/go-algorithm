package main

import (
	"fmt"
)

//递归？？？每次选择都有两个不同的选择 超时了 我去。。可惜。。。。。呜呜呜 
//尝试一下动态规划把

// var ans int  

func main() {

	fmt.Println("hello world")
	m := 3
	n := 2
	ans := 0
	choice(1, 1, &ans, m, n)

	fmt.Printf("ans: %d", ans)

}

func choice(x int, y int, ans *int, m int, n int) {

	if x == m  && y == n {
		//说明到终点了 
		*ans += 1
		// fmt.Println("hello !")
		return 
	}

	if x > m || y > n {
		return 
	}

	choice(x + 1, y, ans, m, n)

	choice(x, y + 1, ans, m, n)

}