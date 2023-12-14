package day11

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day11/input_test")

	r := SolvePart1(input, 1)

	assert.Equal(t, 374, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day11/input_test")

	r := SolvePart1(input, 99)

	assert.Equal(t, 8410, r)
}
