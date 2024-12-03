package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var ans int64 = 0

	disabled := false
	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

		matches := r.FindAllString(line, -1)

		for _, match := range matches {
			var x, y int64
			if match == "don't()" {
				disabled = true
			} else if match == "do()" {
				disabled = false
			} else {
				_, _ = fmt.Sscanf(match, "mul(%d,%d)", &x, &y)

				if !disabled {
					ans += x * y
				}
			}

		}
	}

	fmt.Println(ans)
}
