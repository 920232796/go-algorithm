package main

import (
	"fmt"
)

func main() {
	word1 := "horse"
	word2 := "ros"

	res := minDistance(word1, word2)

	fmt.Println(res)

}

func minDistance(word1 string, word2 string) int {
	word1_len := len(word1) + 1
	word2_len := len(word2) + 1

	dp := make([][]int, word2_len)
	for i := 0; i < word2_len; i ++ {
		dp[i] = make([]int, word1_len)
	}


	for i := 0; i < word1_len; i ++ {
		dp[0][i] = i
	}

	for j := 0; j < word2_len; j ++ {
		dp[j][0] = j
	}

	

	for i := 1; i < word2_len; i ++ {
		for j := 1; j < word1_len; j ++ {

			if word2[i - 1] == word1[j - 1] {
				dp[i][j] = dp[i - 1][j - 1]
			} else {
				dp[i][j] = min(min(dp[i - 1][j] + 1, dp[i][j - 1] + 1), dp[i - 1][j - 1] + 1)
			}
		}
	}

	fmt.Println(dp)

	return dp[word2_len - 1][word1_len - 1]
}

func min(i int, j int) int {

	if i > j {
		return j
	} 

	return i
	
}