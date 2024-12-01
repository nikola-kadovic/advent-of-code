package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		dist += abs(right[i] - left[i])
	}

	fmt.Println(dist)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
