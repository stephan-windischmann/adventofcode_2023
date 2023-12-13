package day10

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day10/input_test1")
	r1 := SolvePart1(input)
	assert.Equal(t, 4, r1)

	input = util.LoadInput("../../../input/day10/input_test2")
	r2 := SolvePart1(input)
	assert.Equal(t, 8, r2)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day10/input_test3")
	r1 := SolvePart2(input)
	assert.Equal(t, 4, r1)

	input = util.LoadInput("../../../input/day10/input_test4")
	r2 := SolvePart2(input)
	assert.Equal(t, 8, r2)

	input = util.LoadInput("../../../input/day10/input_test5")
	r3 := SolvePart2(input)
	assert.Equal(t, 10, r3)
}
