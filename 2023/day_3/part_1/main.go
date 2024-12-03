package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var board []string

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, line)
	}

	var numbers []int

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if unicode.IsDigit(rune(board[i][j])) {
				num := captureNumber(board, i, j)
				if isNearSymbol(board, i, j) {
					numbers = append(numbers, num)
				}
			}
			for j < len(board[0]) && unicode.IsDigit(rune(board[i][j])) {
				j++
			}
		}
	}

	ans := 0

	for _, num := range numbers {
		ans += num
	}

	fmt.Println(ans)
}

func isNearSymbol(board []string, i, j int) bool {
	for ; unicode.IsDigit(rune(board[i][j])) && j < len(board[0]); j++ {
		for _, p := range utils.DiffAdjDiagonal {
			if withinBounds(i+p.First, j+p.Second, len(board), len(board[0])) && isSymbol(board[i+p.First][j+p.Second]) {
				return true
			}
		}
	}

	return false
}

func isSymbol(r byte) bool {
	return !unicode.IsDigit(rune(r)) && r != '.'
}

func withinBounds(i, j, n, m int) bool {
	return i >= 0 && i < n && j >= 0 && j < m
}

func captureNumber(board []string, i, j int) int {
	var nums []int
	for unicode.IsDigit(rune(board[i][j])) {
		nums = append(nums, int(board[i][j]-'0'))
		j++

		if j == len(board[i]) {
			break
		}
	}

	num := 0
	place := 1

	for idx := len(nums) - 1; idx >= 0; idx-- {
		num += nums[idx] * place
		place *= 10
	}

	return num
}
