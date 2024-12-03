package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var strDigits = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ans := 0

	for scanner.Scan() {
		line := scanner.Text()

		var firstDigit, lastDigit int

		for i, val := range line {
			if digit := containsDigitSuffix(line[:i]); digit != -1 {
				firstDigit = digit
				break
			}
			if unicode.IsDigit(val) {
				firstDigit = int(val - '0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if digit := containsDigitPrefix(line[i:]); digit != -1 {
				lastDigit = digit
				break
			}
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = int(line[i] - '0')
				break
			}
		}
		ans += firstDigit*10 + lastDigit
	}

	fmt.Println(ans)
}

func containsDigitSuffix(s string) int {
	for i, digit := range strDigits {
		if strings.HasSuffix(s, digit) {
			return i+1
		}
	}

	return -1
}

func containsDigitPrefix(s string) int {
	for i, digit := range strDigits {
		if strings.HasPrefix(s, digit) {
			return i+1
		}
	}

	return -1
}

