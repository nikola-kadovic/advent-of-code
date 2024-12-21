package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]rune{}

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	m, n := len(grid), len(grid[0])

	explored := utils.Make2DArr[bool](m, n)

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if explored[i][j] {
				continue
			}

			area, perimeter := explore(grid, i, j, explored)

			ans += area * perimeter
		}
	}

	fmt.Println(ans)
}

func explore(grid [][]rune, i, j int, explored [][]bool) (area int, perimeter int) {
	m, n := len(grid), len(grid[0])
	queue := utils.NewQueue[utils.Vec]()
	queue.Enqueue(utils.Vec{X: i, Y: j})

	region := grid[i][j]

	for queue.Len() > 0 {
		front := queue.Dequeue()

		if !front.WithinBounds(m, n) || utils.GetAtVec(front, explored) || utils.GetAtVec(front, grid) != region {
			continue
		}

		utils.SetAtVec(front, explored, true)

		perimeter += calculatePerimeter(grid, front)
		area++

		for _, diff := range utils.DiffAdj {
			queue.Enqueue(front.Add(diff))
		}
	}

	return
}

func calculatePerimeter(grid [][]rune, loc utils.Vec) int {
	m, n := len(grid), len(grid[0])
	perimeter := 0

	for _, d := range utils.DiffAdj {
		if !loc.Add(d).WithinBounds(m, n) || (utils.GetAtVec(loc.Add(d), grid) != utils.GetAtVec(loc, grid)) {
			perimeter++
		}
	}

	return perimeter
}
