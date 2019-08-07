package main

import (
	"fmt"
)

//递归？？？每次选择都有两个不同的选择 超时了 我去。。可惜。。。。。呜呜呜 
//尝试一下动态规划把

// var ans int  

func main() {

	m := 7
	n := 3

	ans := uniquePathss(m, n)

	fmt.Printf("ans: %d", ans)

}

func uniquePathss(m int, n int) int {
	
	dp := make([][]int, m)

	for i := 0; i < m; i ++ {
	
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i ++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j ++ {
		dp[0][j] = 1
	}

	for i := 0; i < m - 1 ; i ++ {
		for j := 0; j < n - 1; j ++ {
			dp[i + 1][j + 1] = dp[i][j + 1] + dp[i + 1][j]
		}
	}

	return dp[m - 1][n - 1]
}
