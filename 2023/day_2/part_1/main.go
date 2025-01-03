package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		isOkay := true

		var header, game string
		header, game, _ = strings.Cut(line, ": ")

		var index int

		fmt.Sscanf(header, "Game %d", &index)

		for _, run := range strings.Split(game, ";") {

			for _, play := range strings.Split(run, ",") {
				var amt int
				var color string

				fmt.Sscanf(play, "%d %s", &amt, &color)

				switch color {
				case "red":
					isOkay = isOkay && amt <= 12
				case "green":
					isOkay = isOkay && amt <= 13
				case "blue":
					isOkay = isOkay && amt <= 14
				}
				if !isOkay {
					break
				}
			}
		}

		ans += index * utils.BoolToInt(isOkay)
	}

	fmt.Println(ans)
}
