package day17

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day17/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 102, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day17/input_test")
	r := SolvePart2(input)
	assert.Equal(t, 94, r)

	input = util.LoadInput("../../../input/day17/input_test2")
	r = SolvePart2(input)
	assert.Equal(t, 71, r)
}
