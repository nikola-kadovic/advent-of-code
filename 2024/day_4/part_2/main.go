package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	board := []string{}

	for scanner.Scan() {
		board = append(board, scanner.Text())
	}

	ans := 0

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			ans += explorePoint(board, i, j)
		}
	}

	fmt.Println(ans)
}

func explorePoint(board []string, i, j int) int {
	if board[i][j] != 'A' {
		return 0
	}

	occurrences := 0

	if (board[i-1][j+1] == 'M' && board[i+1][j-1] == 'S') || (board[i-1][j+1] == 'S' && board[i+1][j-1] == 'M') {
		occurrences++
	}

	if (board[i-1][j-1] == 'M' && board[i+1][j+1] == 'S') || (board[i-1][j-1] == 'S' && board[i+1][j+1] == 'M') {
		occurrences++
	}

	if occurrences == 2 {
		return 1
	}
	return 0
}
