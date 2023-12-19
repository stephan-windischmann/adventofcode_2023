package day12

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day12/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 21, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day12/input_test")

	r := SolvePart2(input)

	assert.Equal(t, 525152, r)
}
