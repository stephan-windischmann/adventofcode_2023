package day3

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day3/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 4361, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day3/input_test")

	r := SolvePart2(input)

	assert.Equal(t, 467835, r)
}
