package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day8/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 6, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day8/input_test2")

	r := SolvePart2(input)

	assert.Equal(t, 6, r)
}
