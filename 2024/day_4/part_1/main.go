package main

import (
	"bufio"
	"fmt"
	"os"
)

var word []byte = []byte{'X', 'M', 'A', 'S'}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	board := []string{}

	for scanner.Scan() {
		board = append(board, scanner.Text())
	}

	ans := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			ans += explorePoint(board, i, j)
			// fmt.Print(explorePoint(board, i, j), " ")
		}
		// fmt.Println()
	}

	fmt.Println(ans)
}

func explorePoint(board []string, i, j int) int {
	if board[i][j] != 'X' {
		return 0
	}

	m := len(board)
	n := len(board[i])

	occurrences := []bool{true, true, true, true, true, true, true, true}
	explored := []int{0, 0, 0, 0, 0, 0, 0, 0}

	for x := 0; x < 4; x++ {
		if j+x < n {
			occurrences[0] = occurrences[0] && (board[i][j+x] == word[x])
			explored[0]++
		}
		if j-x >= 0 {
			occurrences[1] = occurrences[1] && board[i][j-x] == word[x]
			explored[1]++
		}

		if i+x < m {
			occurrences[2] = occurrences[2] && board[i+x][j] == word[x]
			explored[2]++
		}
		if i-x >= 0 {
			occurrences[3] = occurrences[3] && board[i-x][j] == word[x]
			explored[3]++
		}

		if i+x < m && j+x < n {
			occurrences[4] = occurrences[4] && board[i+x][j+x] == word[x]
			explored[4]++
		}
		if i-x >= 0 && j-x >= 0 {
			occurrences[5] = occurrences[5] && board[i-x][j-x] == word[x]
			explored[5]++
		}

		if i+x < m && j-x >= 0 {
			occurrences[6] = occurrences[6] && board[i+x][j-x] == word[x]
			explored[6]++
		}
		if i-x >= 0 && j+x < n {
			occurrences[7] = occurrences[7] && board[i-x][j+x] == word[x]
			explored[7]++
		}
	}

	count := 0

	for i, occurrence := range occurrences {
		if occurrence && explored[i] == 4 {
			count++
		}
	}

	return count
}
