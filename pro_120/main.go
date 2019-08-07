package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello world;")

	triangle := [][]int {{2},
						{3, 4},
						{6, 5, 7},
						{4, 1, 8, 3},	
				}
	res := minimumTotal(triangle)

	fmt.Println(res)
}

func minimumTotal(triangle [][]int) int {
    
    dp := make([][]int, len(triangle))
	
	for i := 0; i < len(triangle); i ++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]
	dp[1][0] = triangle[0][0] + triangle[1][0]
	dp[1][1] = triangle[0][0] + triangle[1][1]

	for j := 2; j < len(triangle); j ++ {
		for k := 0; k < len(dp[j]); k ++ {
			if k - 1 < 0 {
				dp[j][k] = dp[j - 1][k] + triangle[j][k]
				continue
			}

			if k + 1 > len(triangle[j - 1]) {
				//右侧出界
				dp[j][k] = dp[j - 1][k - 1] + triangle[j][k]
				continue
			}
			fmt.Println(j, k)
			dp[j][k] = min(dp[j - 1][k], dp[j - 1][k - 1]) + triangle[j][k]
		}
	}

	fmt.Println(dp)

	return 0
}

func min(i int, j int) int {
	if i > j {
		return j  
	} 

	return i 
}