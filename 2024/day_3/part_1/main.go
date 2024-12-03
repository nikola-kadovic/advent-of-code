package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ans := 0

	for scanner.Scan() {
		line := scanner.Text()

		r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

		matches := r.FindAllString(line, -1)

		for _, match := range matches {
			var x, y int
			_, _ = fmt.Sscanf(match, "mul(%d,%d)", &x, &y)

			ans += x * y
		}
	}

	fmt.Println(ans)
}
