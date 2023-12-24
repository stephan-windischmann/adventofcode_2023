package day19

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day19/input_test")

	r := SolvePart1(input)

	assert.Equal(t, 19114, r)
}

func TestSolvePart2(t *testing.T) {
	input := util.LoadInput("../../../input/day19/input_test")

	r := SolvePart2(input)

	assert.Equal(t, 167409079868000, r)
}
