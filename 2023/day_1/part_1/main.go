package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ans := 0

	for scanner.Scan() {
		line := scanner.Text()

		var firstDigit, lastDigit int

		for _, val := range line {
			if unicode.IsDigit(val) {
				firstDigit = int(val - '0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = int(line[i] - '0')
				break
			}
		}
		ans += firstDigit*10 + lastDigit
	}

	fmt.Println(ans)
}
