package day20

import (
	"testing"

	"github.com/stephan-windischmann/adventofcode_2023/internal/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	input := util.LoadInput("../../../input/day20/input_test1")
	r := SolvePart1(input)
	assert.Equal(t, 32000000, r)

	input = util.LoadInput("../../../input/day20/input_test2")
	r = SolvePart1(input)
	assert.Equal(t, 11687500, r)
}
