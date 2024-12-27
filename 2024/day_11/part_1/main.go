package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	arr := []int{}

	for _, strVal := range strings.Split(scanner.Text(), " ") {
		val, _ := strconv.Atoi(strVal)
		arr = append(arr, val)
	}

	for i := 0; i < 25; i++ {
		arr = update(arr)
	}

	fmt.Println(len(arr))
}

func update(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[i] = 1
		} else if numLen(arr[i])%2 == 0 {
			l, r := split(arr[i])
			arr[i] = l
			arr = slices.Insert(arr, i+1, r)
			i++
		} else {
			arr[i] *= 2024
		}
	}

	return arr
}

func split(num int) (int, int) {
	if numLen(num)%2 != 0 {
		panic("num must be even")
	}

	l := numLen(num)

	x, y := 0, 0

	place := 1
	for i := 0; i < l/2; i++ {
		y += (num % 10) * place
		num /= 10
		place *= 10
	}

	place = 1
	for i := 0; i < l/2; i++ {
		x += (num % 10) * place
		num /= 10
		place *= 10
	}

	return x, y
}

func numLen(num int) int {
	l := 0
	for num > 0 {
		num /= 10
		l++
	}

	return l
}
