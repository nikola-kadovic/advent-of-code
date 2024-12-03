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

	var numTickets map[int]int = map[int]int{}
	ans := 0
	idx := 0

	for scanner.Scan() {
		line := scanner.Text()
		score := getScore(line)
		for i := idx + 1; i < idx+1+score; i++ {
			numTickets[i] += numTickets[idx] + 1
		}

		ans += 1 + numTickets[idx]

		idx++
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
			score++
		}
	}

	return score
}
