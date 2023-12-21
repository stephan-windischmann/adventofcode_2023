package day14

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day14/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 136, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day14/input_test")

	r := SolvePart2(input)

	assert.Equal(t, 64, r)
}
