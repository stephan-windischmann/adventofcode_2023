package util

import (
	"bufio"
	"os"
)

func LoadInput(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var input []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func SumInt(in []int) int {
	sum := 0

	for _, n := range in {
		sum += n
	}

	return sum
}

func ProdInt(in []int) int {
	sum := in[0]

	for i := 1; i < len(in); i++ {
		sum *= in[i]
	}

	return sum
}

func CreateRuneTable(input []string) [][]rune {
	t := make([][]rune, 0, len(input))

	for _, l := range input {
		t = append(t, []rune(l))
	}

	return t
}

func RuneToInt(c rune) int {
	return int(c - '0')
}
