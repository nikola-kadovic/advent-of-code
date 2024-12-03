package utils

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

type Pair struct {
	First, Second int
}

var DiffAdj = []Pair{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var DiffAdjDiagonal []Pair = []Pair{
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
