package main


import (
	"fmt"
)

func main() {

	fmt.Println("hello world;")
	matrix := [][] int {{0, 1, 2, 0}, 
						{3, 4, 5, 2},
						{1, 3, 1, 5},			
					}
	setZeroes(matrix)
}

func setZeroes(matrix [][]int)  {
	fmt.Println(matrix)
	row := []int {}
	col := []int {}

	for i := 0; i < len(matrix); i ++ {
		for j := 0; j < len(matrix[0]); j ++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}

	for _, v := range(row) {
		for j := 0; j < len(matrix[v]); j ++ {
			matrix[v][j] = 0
		}
	}

	for _, v := range(col) {
		for j := 0; j < len(matrix); j ++ {
			matrix[j][v] = 0
		}
	}


	// fmt.Println(row, col)

	fmt.Println(matrix)
}