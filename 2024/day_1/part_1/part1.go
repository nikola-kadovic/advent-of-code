package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

func main() {
	var left, right []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		var l, r int

		fmt.Sscanf(line, "%d %d", &l, &r)

		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	dist := 0

	for i := range len(left) {
		dist += utils.Abs(right[i] - left[i])
	}

	fmt.Println(dist)
}
