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

	solve(0, array)
	
	return nil
}

func solve(row int, array []int ) {

	if row == len(array) {
		//证明找到解了
		fmt.Println(array)
		return 
	}
	
	for i := 0; i < len(array); i ++ { //循环每一列
		array[row] = i 
		if justi(row, i, array) {
			//判断是否有正确的解 如果有的话就返回true 那么就递归去下一行
			solve(row + 1, array)
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

