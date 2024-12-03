package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var left, right []int

	scanner := bufio.NewScanner(os.Stdin)

	var mp map[int]int = make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		var l, r int

		fmt.Sscanf(line, "%d %d", &l, &r)

		left = append(left, l)
		right = append(right, r)
		mp[r]++
	}

	score := 0

	for _, val := range left {
		score += val * mp[val]
	}

	fmt.Println(score)
}
