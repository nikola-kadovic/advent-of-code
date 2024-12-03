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