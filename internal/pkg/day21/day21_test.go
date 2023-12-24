package day21

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day21/input_test")

	r := SolvePart1(input, 6)

	assert.Equal(t, 16, r)
}
