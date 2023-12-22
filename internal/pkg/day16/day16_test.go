package day16

import (
	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day16/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 46, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day16/input_test")

	r := SolvePart2(input)

	assert.Equal(t, 51, r)
}
