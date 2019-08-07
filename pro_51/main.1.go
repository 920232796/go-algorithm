package main

import (
	"fmt"
	"math"
)

//回溯 + 搜索
func main() {

	fmt.Println("hello world;")

	res := solveNQueens(4)
	fmt.Println(res)

}

func solveNQueens(n int) [][]string {
	
	array := make([]int, n)
	resInt := [][]int {}
	res := [][]string{}
	solve(&resInt, 0, array)
	fmt.Printf("resInt: %v\n", resInt)
	row := ""
	each := []string {}
	for i := 0; i < n; i ++ {
		row += "."
	}
	for j := 0; j < n; j ++ {
		each = append(each, row)
	}

	for i := 0; i < len(resInt); i ++ {
		for j := 0; j < n; j ++ {
			res[i][j][resInt[i][j]] = "Q"
		}
	}
	return res
}

func solve(resInt *[][]int, row int, array []int ) {

	if row == len(array) {
		//证明找到解了
		temp := make([]int, len(array))
		copy(temp, array)
		fmt.Println(array)
		*resInt = append(*resInt, temp)
		return 
	}
	
	for i := 0; i < len(array); i ++ { //循环每一列
		array[row] = i 
		if justi(row, i, array) {
			//判断是否有正确的解 如果有的话就返回true 那么就递归去下一行
			solve(resInt, row + 1, array)
		}
		//如果不合法的话 array 重新置为0
	}

	
}


func justi(row int, col int, array []int ) bool {

	for i := 0; i < row; i ++ {
		//循环前面的行 判断是否合法
		if array[i] == col || math.Abs(float64(row - i)) == math.Abs(float64(col - array[i])) {
			//证明位于同一列或者同一斜线
			return false
		}
		
	}
	return true 
}

