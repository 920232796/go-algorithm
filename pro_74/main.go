package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello world;")
	matrix := [][]int {
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 13

	res := searchMatrix(matrix, target)

	fmt.Println(res)
}

func searchMatrix(matrix [][]int, target int) bool {
	
	for i := 0; i < len(matrix); i ++ {
		if target < matrix[i][len(matrix[0]) - 1] {
			//搜索那一行就行了
			for j := 0; j < len(matrix[0]); j ++ {
				if target == matrix[i][j] {
					return true
				}
			}
			return false
		}
	} 
	return false
}