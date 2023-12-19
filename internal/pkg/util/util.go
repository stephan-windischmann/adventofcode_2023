package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

func StringToIntSlice(in string) []int {
	nums := make([]int, 0)

	for _, n := range strings.Fields(in) {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	return nums
}

func AllEquals(in []int, eq int) bool {
	for _, n := range in {
		if n != eq {
			return false
		}
	}
	return true
}

func ToRune(input []string) [][]rune {
	runes := make([][]rune, len(input))

	for i := 0; i < len(input); i++ {
		l := make([]rune, len(input[i]))
		for j := 0; j < len(input[i]); j++ {
			l[j] = rune(input[i][j])
		}
		runes[i] = l
	}

	return runes
}

func JoinIntList(input []int, sep string) string {
	sList := make([]string, len(input))

	for i, n := range input {
		sList[i] = strconv.Itoa(n)
	}

	return strings.Join(sList, sep)
}
