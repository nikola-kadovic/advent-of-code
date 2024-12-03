package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		ans += getScore(line)
	}

	fmt.Println(ans)
}

func getScore(line string) int {
	score := 0

	_, rest, _ := strings.Cut(line, ": ")
	strCardsNeed, strCardsHave, _ := strings.Cut(rest, " | ")

	cardsNeed := map[int]bool{}

	for _, card := range strings.Fields(strCardsNeed) {
		c, _ := strconv.Atoi(card)
		cardsNeed[c] = true
	}

	for _, card := range strings.Fields(strCardsHave) {
		c, _ := strconv.Atoi(card)
		if cardsNeed[c] {
			if score == 0 {
				score++
			} else {
				score *= 2
			}
		}
	}

	return score
}
