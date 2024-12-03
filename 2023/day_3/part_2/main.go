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

	ans := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '*' {
				var adjNums []int
				var covered *map[int]map[int]bool = &map[int]map[int]bool{}
				for _, d := range utils.DiffAdjDiagonal {
					if utils.GetFrom2DMap(covered, i+d.First, j+d.Second) {
						continue
					}

					if num := captureNumber(covered, board, i+d.First, j+d.Second); num != -1 {
						adjNums = append(adjNums, num)
					}
				}

				if len(adjNums) == 2 {
					ans += adjNums[0] * adjNums[1]
				}
			}
		}
	}

	fmt.Println(ans)
}

func withinBounds(i, j, n, m int) bool {
	return i >= 0 && i < n && j >= 0 && j < m
}

func captureNumber(covered *map[int]map[int]bool, board []string, i, j int) int {
	if !withinBounds(i, j, len(board), len(board[0])) {
		return -1
	}
	if !unicode.IsDigit(rune(board[i][j])) {
		return -1
	}

	var numsBefore, numsAfter []int
	jBefore, jAfter := j-1, j
	for jAfter < len(board[i]) && unicode.IsDigit(rune(board[i][jAfter])) {
		(*covered)[i][jAfter] = true
		utils.SetFrom2DMap(covered, i, jAfter, true)
		numsAfter = append(numsAfter, int(board[i][jAfter]-'0'))
		jAfter++
	}

	for jBefore >= 0 && unicode.IsDigit(rune(board[i][jBefore])) {
		utils.SetFrom2DMap(covered, i, jBefore, true)
		numsBefore = append(numsBefore, int(board[i][jBefore]-'0'))
		jBefore--
	}

	num := 0
	place := 1

	for idx := len(numsAfter) - 1; idx >= 0; idx-- {
		num += numsAfter[idx] * place
		place *= 10
	}

	for idx := 0; idx < len(numsBefore); idx++ {
		num += numsBefore[idx] * place
		place *= 10
	}

	return num
}
