package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

type vec struct {
	x, y int
}

func (v vec) withinBounds(m, n int) bool {
	return v.x >= 0 && v.x < m && v.y >= 0 && v.y < n
}

func (v vec) add(other vec) vec {
	return vec{x: v.x + other.x, y: v.y + other.y}
}

func dist(a, b vec) vec {
	return vec{x: a.x - b.x, y: a.y - b.y}
}

func setAtVec[T any](input [][]T, v vec, val T) {
	input[v.x][v.y] = val
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]rune{}

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	m, n := len(grid), len(grid[0])

	containsAntiNode := utils.Make2DArr[bool](m, n)

	antennaLoc := map[rune][]vec{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' {
				antennaLoc[grid[i][j]] = append(antennaLoc[grid[i][j]], vec{i, j})
			}
		}
	}

	for _, antennas := range antennaLoc {
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				if antennas[i].add(dist(antennas[i], antennas[j])).withinBounds(m, n) {
					setAtVec(containsAntiNode, antennas[i].add(dist(antennas[i], antennas[j])), true)
				}
				if antennas[j].add(dist(antennas[j], antennas[i])).withinBounds(m, n) {
					setAtVec(containsAntiNode, antennas[j].add(dist(antennas[j], antennas[i])), true)
				}
			}
		}
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if containsAntiNode[i][j] {
				ans++
			}
		}
	}

	fmt.Println(ans)
}
