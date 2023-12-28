package day24

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day24/input_test")

	r := SolvePart1(input, 7, 27)

	assert.Equal(t, 2, r)
}
