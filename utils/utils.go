package utils

import (
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

type Vec struct {
	First, Second int
}

func (v Vec) Add(other Vec) Vec {
	return Vec{v.First + other.First, v.Second + other.Second}
}

var DiffAdj = []Vec{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var DiffAdjDiagonal []Vec = []Vec{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func GetFrom2DMap[T comparable, V any](mp *map[T]map[T]V, first, second T) V {
	if _, ok := (*mp)[first]; !ok {
		(*mp)[first] = make(map[T]V)
	}

	return (*mp)[first][second]
}

func SetFrom2DMap[T comparable, V any](mp *map[T]map[T]V, first, second T, val V) {
	if _, ok := (*mp)[first]; !ok {
		(*mp)[first] = make(map[T]V)
	}

	(*mp)[first][second] = val
}

func Pow[T Number](a, b T) T {
	if b == 0 {
		return 1
	}

	if b%2 == 0 {
		return Pow(a*a, b/2)
	}

	return a * Pow(a, b-1)
}

func StrArrToIntArr(arr, delim string) []int {
	res := []int{}

	for _, e := range strings.Split(arr, delim) {
		val, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}

		res = append(res, val)
	}

	return res
}

func Make2DArr[T any](m, n int) [][]T {
	arr := make([][]T, m)
	for i := range arr {
		arr[i] = make([]T, n)
	}

	return arr
}

func CopyMap[K comparable, V any](mp map[K]V) map[K]V {
	newMap := map[K]V{}

	for k, v := range mp {
		newMap[k] = v
	}

	return newMap
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return Gcd(b, a%b)
}
