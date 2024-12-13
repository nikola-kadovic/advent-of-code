package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type pos struct {
	loc, dir vec
}

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

	m, n := len(input), len(input[0])

	var wg sync.WaitGroup = sync.WaitGroup{}
	outputs := make(chan int, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if input[i][j] == '#' || input[i][j] == '^' {
				continue
			}
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				outputs <- simulatePathWithBarrier(input, vec{startX, startY}, vec{i, j})
			}(i, j)
		}
	}

	wg.Wait()
	close(outputs)

	ans := 0

	for val := range outputs {
		ans += val
	}

	fmt.Println(ans)
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

func simulatePathWithBarrier(input [][]rune, start, barrier vec) int {
	m, n := len(input), len(input[0])
	visited := map[pos]bool{}

	loc := start
	// up
	dir := vec{-1, 0}

	ans := 0

	for loc.inBounds(m, n) {
		if visited[pos{loc, dir}] {
			return 1
		}

		visited[pos{loc, dir}] = true

		for loc.add(dir).inBounds(m, n) && (getAtVec(input, loc.add(dir)) == '#' || loc.add(dir) == barrier) {
			dir = dir.rotate()
			visited[pos{loc, dir}] = true
		}

		loc = loc.add(dir)
	}

	return ans
}
