package main

func solveSudoku(board [][]byte) {

	var horizontal [9][9]bool
	var vertical [9][9]bool
	var matrix [9][9]bool

	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				v := board[i][j] - '1'
				horizontal[i][v] = true
				vertical[j][v] = true
				matrix[(i/3)*3+j/3][v] = true
			}
		}
	}

	dfs3(horizontal, vertical, matrix, board, 0, 0)

}

func dfs3(horizontal [9][9]bool, vertical [9][9]bool, matrix [9][9]bool, board [][]byte, i int, j int) bool {
	for board[i][j] != '.' {
		j += 1
		if j >= 9 {
			j = 0
			i += 1
		}
		if i >= 9 {
			return true
		}
	}

	for n := 0; n < 9; n++ {
		if !horizontal[i][n] && !vertical[j][n] && !matrix[(i/3)*3+j/3][n] {
			horizontal[i][n] = true
			vertical[j][n] = true
			matrix[(i/3)*3+j/3][n] = true
			board[i][j] = byte(n + '1')
			if dfs3(horizontal, vertical, matrix, board, i, j) {
				return true
			} else {
				horizontal[i][n] = false
				vertical[j][n] = false
				matrix[(i/3)*3+j/3][n] = false
				board[i][j] = '.'
			}
		}
	}

	return false
}
