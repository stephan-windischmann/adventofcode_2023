package day6

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day6/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 288, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day6/input_test")

	r := SolvePart2(input)

	assert.Equal(t, 71503, r)
}
