package main

import (
	"fmt"
)

func main() {

	obstacleGrid := [][]int {
		{0, 0},
		{1, 1},
		{0, 0},
		
	}

	res := uniquePathsWithObstacles(obstacleGrid)

	fmt.Println(res)

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    
    m := len(obstacleGrid[0])
    n := len(obstacleGrid)
    
    dp := make([][]int, n)
    
    for i := 0; i < n; i ++ {
        dp[i] = make([]int, m)
    }
    if obstacleGrid[0][0] == 0 {
        dp[0][0] = 1
    } else {
        dp[0][0] = 0
        return 0
    }
    
    flag := 0
    for i := 1; i < m; i ++ {
        if obstacleGrid[0][i] == 0 && flag == 0 {
            dp[0][i] = 1
        } else  {
            flag = 1
            dp[0][i] = 0
        }
    }
    flag = 0
    for j := 1; j < n; j ++ {
        if obstacleGrid[j][0] == 0 && flag == 0 {
            dp[j][0] = 1
        } else {
            flag = 1
            dp[j][0] = 0
        }
	}
	

	for _, v := range(dp) {
		fmt.Println(v)
	}
    
    for i := 0; i < n - 1; i ++ {
        for j := 0; j < m - 1; j ++ {
            
            if obstacleGrid[i + 1][j + 1] == 0 {
                dp[i + 1][j + 1] = dp[i + 1][j] + dp[i][j + 1]
            } else {
                dp[i + 1][j + 1] = 0
            }
           
        }
	}
	
	for _, v := range(dp) {
		fmt.Println(v)
	}
    
    return dp[n - 1][m - 1]
    
}