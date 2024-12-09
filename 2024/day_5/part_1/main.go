package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	orders := []utils.Pair{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		var a, b int
		fmt.Sscanf(line, "%d|%d", &a, &b)

		orders = append(orders, utils.Pair{First: a, Second: b})
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
			_, okF := pos[order.First]
			_, okS := pos[order.Second]

			if !okF || !okS {
				continue
			}

			if pos[order.First] > pos[order.Second] {
				inOrder = false
				break
			}
		}

		if inOrder {
			ans += pages[len(pages)/2]
		}
	}

	fmt.Println(ans)
}
