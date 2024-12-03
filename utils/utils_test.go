package utils_test

import (
	"testing"

	"github.com/nikola-kadovic/advent-of-code/utils"
	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	tests := []struct {
		base, exp, want int
	}{
		{2, 3, 8},
		{3, 2, 9},
		{5, 0, 1},
		{20, 4, 160000},
	}

	for _, tt := range tests {
		got := utils.Pow(tt.base, tt.exp)
		assert.Equal(t, tt.want, got)
	}
}
