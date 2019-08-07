package main

import (
	"fmt"
)

//递归
func main() {

	fmt.Println("hello world;")

	board := [][] byte {
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	// board := [][] byte {
	// 	{'C', 'A', 'A'},
	// 	{'A', 'A', 'A'},
	// 	{'B', 'C', 'D'},
	// }
	word := "ABCCSC"
	res := exist(board, word)
	fmt.Println(res)
}

func exist(board [][]byte, word string) bool {

	fmt.Println(board)
	flag := 0
	
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j ++ {
			if board[i][j] == word[0] {
				dfs(&flag, board, i, j, word, 0)
				if flag == 1 {
					return true 
				} 
			}
			
		}
	}
	
	return false
}

func dfs(flag *int, board [][]byte, i int, j int, word string, index int)  {

	if index == len(word) {
		*flag = 1
		return 
	}

	if i < 0 || j < 0 || i > len(board) - 1 || j > len(board[0]) - 1 {
		return 
	}

	if board[i][j] != word[index] {
		return

	}

	if board[i][j] == word[index] {

		println(i, j, index)
		temp := board[i][j]
		// fmt.Printf("board 左移前: %d", board[i][j])
		board[i][j] = '1'
		// fmt.Printf("board 左移后: %d", board[i][j])
		dfs(flag, board, i, j + 1, word, index + 1)
		dfs(flag, board, i + 1, j, word, index + 1)
		dfs(flag, board, i - 1, j, word, index + 1)
		dfs(flag, board, i, j - 1, word, index + 1)
		// board[i][j] >>= 2
		board[i][j] = temp 
		
	}
}