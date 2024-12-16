package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type testcase struct {
	ans  int
	nums []int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	testcases := []testcase{}

	for scanner.Scan() {
		ansStr, numsStr, _ := strings.Cut(scanner.Text(), ": ")

		ans, _ := strconv.Atoi(ansStr)
		nums := []int{}

		for _, numStr := range strings.Split(numsStr, " ") {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}

		testcases = append(testcases, testcase{ans, nums})
	}

	ans := 0

	for _, tc := range testcases {
		ans += solve(tc)
	}

	fmt.Println(ans)
}

func solve(tc testcase) int {
	if dfs(tc.ans, tc.nums[0], tc.nums[1:]) {
		return tc.ans
	}

	return 0
}

func dfs(want, have int, nums []int) bool {
	if len(nums) == 0 {
		return have == want
	}
	// try to multiply
	if dfs(want, have*nums[0], nums[1:]) {
		return true
	}

	// try to add
	if dfs(want, have+nums[0], nums[1:]) {
		return true
	}

	// try to concat
	if dfs(want, concat(have, nums[0]), nums[1:]) {
		return true
	}

	// nothing worked
	return false
}

func concat(a, b int) int {
	l := b

	for l > 0 {
		l /= 10
		a *= 10
	}

	return a + b
}
