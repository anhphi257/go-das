package go_das

import (
	"fmt"
	"testing"
)

func Test_LC289(t *testing.T) {
	board := [][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	}
	gameOfLife(board)
	fmt.Println(board)
}
