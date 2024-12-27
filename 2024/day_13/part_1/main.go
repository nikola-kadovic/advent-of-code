package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

type vec struct {
	x, y int
}

type game struct {
	a     vec
	b     vec
	prize vec
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	games := []game{}

	for scanner.Scan() {
		var aX, aY, bX, bY, pX, pY int

		line := scanner.Text()
		fmt.Sscanf(line, "Button A: X+%d, Y+%d", &aX, &aY)

		scanner.Scan()
		line = scanner.Text()
		fmt.Sscanf(line, "Button B: X+%d, Y+%d", &bX, &bY)

		scanner.Scan()
		line = scanner.Text()
		fmt.Sscanf(line, "Prize: X=%d, Y=%d", &pX, &pY)

		scanner.Scan()

		games = append(games, game{vec{aX, aY}, vec{bX, bY}, vec{pX, pY}})
	}

	ans := 0

	for _, g := range games {
		ans += minTokens(g)
	}

	fmt.Println(ans)
}

func minTokens(g game) int {
	minTokens := math.MaxInt

	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			if a*g.a.x+b*g.b.x == g.prize.x && a*g.a.y+b*g.b.y == g.prize.y {
				minTokens = utils.Min(minTokens, (3*a)+b)
			}
		}
	}

	if minTokens == math.MaxInt {
		return 0
	}

	return minTokens
}
