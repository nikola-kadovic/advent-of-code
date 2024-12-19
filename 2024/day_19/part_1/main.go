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
	for _, val := range strings.Split(scanner.Text(), " ") {
		substrings = append(substrings, val)
	}

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
	dp := make([]bool, len(line)+1)
	dp[0] = true

	for i := 0; i <= len(line); i++ {
		for _, sub := range substrings {
			if i+len(sub) <= len(line) && line[i:i+len(sub)] == sub {
				dp[i+len(sub)] = dp[i+len(sub)] || dp[i]
			}
		}
	}

	fmt.Println(dp)

	return dp[len(line)]
}
