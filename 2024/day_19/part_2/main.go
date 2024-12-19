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
		ans += canMakeUpWord(substrings, line)
	}

	fmt.Println(ans)
}

func canMakeUpWord(substrings []string, line string) int {
	// numOccurrences[i] = number of ways we can make up line[:i] with substrings
	numOccurrences := make([]int, len(line)+1)
	numOccurrences[0] = 1

	for i := 0; i < len(line); i++ {
		for _, sub := range substrings {
			if i+len(sub) <= len(line) && line[i:i+len(sub)] == sub {
				numOccurrences[i+len(sub)] += numOccurrences[i]
			}
		}
	}

	return numOccurrences[len(line)]
}
