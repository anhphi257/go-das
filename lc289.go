package go_das

func gameOfLife(board [][]int) {
	n := len(board)
	m := len(board[0])
	next := make([][]int, n)
	for i := range next {
		next[i] = make([]int, m)
	}
	dx := []int{-1, 0, 1}
	dy := []int{-1, 0, 1}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cnt := 0
			for _, ii := range dx {
				for _, jj := range dy {
					if ii == 0 && jj == 0 {
						continue
					}
					if 0 <= i+ii && i+ii < n && 0 <= j+jj && j+jj < m {
						cnt += board[i+ii][j+jj]
					}
				}
			}
			if board[i][j] == 0 {
				if cnt == 3 {
					next[i][j] = 1
				} else {
					next[i][j] = 0
				}
			} else {
				if cnt == 2 || cnt == 3 {
					next[i][j] = 1
				} else {
					next[i][j] = 0
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			board[i][j] = next[i][j]
		}
	}
}
