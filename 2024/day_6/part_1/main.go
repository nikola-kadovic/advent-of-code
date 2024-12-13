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

func (v vec) add(other vec) vec {
	return vec{v.x + other.x, v.y + other.y}
}

func (v vec) rotate() vec {
	return vec{v.y, -v.x}
}

func (v vec) inBounds(m, n int) bool {
	return v.x >= 0 && v.x < m && v.y >= 0 && v.y < n
}

func getAtVec[T any](input [][]T, v vec) T {
	return input[v.x][v.y]
}

func setAtVec[T any](input [][]T, v vec, val T) {
	input[v.x][v.y] = val
}

func main() {
	input := getInput()

	var startX, startY int

	for i, row := range input {
		for j, val := range row {
			if val == '^' {
				startX = i
				startY = j
			}
		}
	}

	fmt.Println(simulatePath(input, startX, startY))
}

func getInput() [][]rune {
	input := [][]rune{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		input = append(input, line)
	}

	return input
}

func simulatePath(input [][]rune, x, y int) int {
	m, n := len(input), len(input[0])
	visited := utils.Make2DArr[bool](m, n)

	pos := vec{x, y}
	// up
	dir := vec{-1, 0}

	ans := 0

	for pos.inBounds(m, n) {
		if !getAtVec(visited, pos) {
			setAtVec(visited, pos, true)
			ans++
		}

		for pos.add(dir).inBounds(m, n) && getAtVec(input, pos.add(dir)) == '#' {
			dir = dir.rotate()
		}

		pos = pos.add(dir)
	}

	return ans
}
