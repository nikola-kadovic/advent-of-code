package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nikola-kadovic/advent-of-code/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var a [][]int

	for scanner.Scan() {
		line := scanner.Text()

		var tmp []int

		for _, val := range strings.Split(line, " ") {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			tmp = append(tmp, num)
		}
		a = append(a, tmp)
	}

	count := 0

	for _, row := range a {
		if isOkay(row) {
			count++
			continue
		}

		for toRemove := 0; toRemove < len(row); toRemove++ {
			dampened := append(append([]int{}, row[:toRemove]...), row[toRemove+1:]...)
			if isOkay(dampened) {
				count++
				break
			}
		}
	}

	fmt.Println(count)
}

func isOkay(row []int) bool {
	isDecreasing := row[0] > row[1]

	for i := 0; i < len(row)-1; i++ {
		if (isDecreasing && row[i] < row[i+1]) || (!isDecreasing && row[i] > row[i+1]) {
			return false
		}
		if utils.Abs(row[i]-row[i+1]) < 1 || utils.Abs(row[i]-row[i+1]) > 3 {
			return false
		}
	}
	return true
}
