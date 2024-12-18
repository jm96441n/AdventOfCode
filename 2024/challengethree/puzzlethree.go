package challengethree

import (
	"AdventOfCode/utils"
	"errors"
	"fmt"
	"strconv"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
	input := utils.OpenFileIntoString(filename)
	p1 := partOne(input)
	p2 := partTwo(input)
	return Result{
		PartOne: p1,
		PartTwo: p2,
	}, nil
}

func partOne(input string) int {
	i := 0
	sum := 0
	res := 0
	for i < len(input) {
		res, i = checkMul(input, i)
		sum += res
		i++
	}
	return sum
}

func checkMul(input string, i int) (int, int) {
	if i+4 < len(input) && input[i:i+4] == "mul(" {
		i += 4
		v1 := -1
		v2 := -1
		v1, i, err := extractDigit(input, i, ',')
		if err != nil {
			return 0, i - 1
		}

		i += 1
		v2, i, err = extractDigit(input, i, ')')
		if err != nil {
			return 0, i - 1
		}

		return v1 * v2, i
	}
	return 0, i
}

func partTwo(input string) int {
	i := 0
	sum := 0
	res := 0
	enabled := true
	for i < len(input) {
		if enabled {
			if i+7 < len(input) && input[i:i+7] == "don't()" {
				enabled = false
				i += 7
				continue
			}
			res, i = checkMul(input, i)
			sum += res
		} else {
			if i+4 < len(input) && input[i:i+4] == "do()" {
				enabled = true
				i += 4
				continue
			}
		}
		i++
	}
	return sum
}

func extractDigit(input string, i int, endChar byte) (int, int, error) {
	var err error
	v := -1
	j := i + 1
	t := 0
	for j < i+4 {
		t, err = strconv.Atoi(input[i:j])
		if err != nil {
			break
		}
		v = t
		j += 1
	}

	i = (j - 1)
	err = nil

	if input[i] != endChar {
		return -1, i, errors.New("wrong end type")
	}
	return v, i, nil
}
