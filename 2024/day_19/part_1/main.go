package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	substrings := []string{}

	scanner.Scan()
	substrings = append(substrings, strings.Split(scanner.Text(), ", ")...)
	scanner.Scan()

	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		if canMakeUpWord(substrings, line) {
			ans++
		}
	}
	fmt.Println(ans)
}

func canMakeUpWord(substrings []string, line string) bool {
	// possible[i] == true iff line[:i] can be made up of substrings
	possible := make([]bool, len(line)+1)
	possible[0] = true

	for i := 0; i < len(line); i++ {
		for _, sub := range substrings {
			if i+len(sub) <= len(line) && line[i:i+len(sub)] == sub {
				possible[i+len(sub)] = possible[i+len(sub)] || possible[i]
			}
		}
	}

	return possible[len(line)]
}
