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

		var minR, minG, minB int = 0, 0, 0

		for _, run := range strings.Split(game, ";") {

			for _, play := range strings.Split(run, ",") {
				var amt int
				var color string

				fmt.Sscanf(play, "%d %s", &amt, &color)

				switch color {
				case "red":
					minR = utils.Max(minR, amt)
				case "green":
					minG = utils.Max(minG, amt)
				case "blue":
					minB = utils.Max(minB, amt)
				}
				if !isOkay {
					break
				}
			}
		}

		ans += minB * minG * minR
	}

	fmt.Println(ans)
}
