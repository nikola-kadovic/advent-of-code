package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	orders := []utils.Vec{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		var a, b int
		fmt.Sscanf(line, "%d|%d", &a, &b)

		orders = append(orders, utils.Vec{X: a, Y: b})
	}

	ans := 0

	for scanner.Scan() {
		pages := utils.StrArrToIntArr(scanner.Text(), ",")

		var pos map[int]int = make(map[int]int)

		for i, page := range pages {
			pos[page] = i
		}

		inOrder := true

		for _, order := range orders {
			_, okF := pos[order.X]
			_, okS := pos[order.Y]

			if !okF || !okS {
				continue
			}

			if pos[order.X] > pos[order.Y] {
				inOrder = false
				break
			}
		}

		if inOrder {
			continue
		}

		var lessThan map[int][]int = make(map[int][]int)

		for _, order := range orders {
			_, okF := pos[order.X]
			_, okS := pos[order.Y]

			if okF && okS {
				lessThan[order.Y] = append(lessThan[order.Y], order.X)
			}
		}

		sorted := []int{}
		done := map[int]bool{}

		for _, page := range pages {
			sorted = addDependencies(sorted, lessThan, page, done)
		}

		ans += sorted[len(sorted)/2]
	}

	fmt.Println(ans)
}

func addDependencies(sortedPages []int, lessThan map[int][]int, page int, done map[int]bool) []int {
	if done[page] {
		return sortedPages
	}

	for _, dep := range lessThan[page] {
		if done[dep] {
			continue
		}

		sortedPages = addDependencies(sortedPages, lessThan, dep, done)
		done[dep] = true
	}

	sortedPages = append(sortedPages, page)
	done[page] = true

	return sortedPages
}
