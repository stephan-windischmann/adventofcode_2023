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

func CountChars(s string) map[rune]int {
	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	return m
}

func CountCharsExclude(s string, exclude rune) map[rune]int {
	m := make(map[rune]int)

	for _, c := range s {
		if c == exclude {
			continue
		}
		m[c]++
	}

	return m
}

func IndexByte(needle byte, l []byte) int {
	for i, b := range l {
		if b == needle {
			return i
		}
	}
	return -1
}
