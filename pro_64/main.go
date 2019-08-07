package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello world;")

	grid := [][]int {
		{1, 3, 1},
		{1, 5, 1}, 
		{4, 2, 1},
	}

	res := minPathSum(grid)

	fmt.Println(res)
}

func minPathSum(grid [][]int) int {
	
	m := len(grid[0])
	n := len(grid)

	dp := make([][]int, n)

	for i := 0; i < n; i ++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < n; i ++ {
		dp[i][0] = grid[i][0] + dp[i - 1][0]
	}

	for j := 1; j < m; j ++ {
		dp[0][j] = grid[0][j]  + dp[0][j - 1]
	}

	for i := 0; i < n - 1; i ++ {

		for j := 0; j < m - 1; j ++ {
			
			temp1 := dp[i + 1][j] + grid[i + 1][j + 1]
			temp2 := dp[i][j + 1] + grid[i + 1][j + 1]

			if temp1 > temp2 {
				dp[i + 1][j + 1] = temp2
			} else {
				dp[i + 1][j + 1] = temp1
			}
		}
	}

	fmt.Println(dp)

	return dp[n - 1][m - 1]

}

