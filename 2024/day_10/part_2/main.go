package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

var diff = []utils.Vec{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	mp := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, val := range line {
			row = append(row, int(val-'0'))
		}
		mp = append(mp, row)
	}

	m, n := len(mp), len(mp[0])

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == 0 {
				ans += calculateScore(i, j, 0, mp)
			}
		}
	}

	fmt.Println(ans)
}

func calculateScore(i, j, h int, mp [][]int) int {
	m, n := len(mp), len(mp[0])

	if mp[i][j] == 9 {
		return 1
	}

	score := 0

	for _, d := range utils.DiffAdj {
		if (utils.Vec{X: i, Y: j}).Add(d).WithinBounds(m, n) && mp[i+d.X][j+d.Y] == h+1 {
			score += calculateScore(i+d.X, j+d.Y, h+1, mp)
		}
	}

	return score
}
