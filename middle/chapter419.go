package middle

func countBattleships(board [][]byte) int {
	ans := 0

	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
	inner:
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' {
				if (i > 0 && board[i-1][j] == 'X') || (j > 0 && board[i][j-1] == 'X') {
					continue inner
				}
				ans++
			}
		}
	}
	return ans
}
