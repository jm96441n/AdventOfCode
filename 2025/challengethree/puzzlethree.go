package challengethree

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode/utils"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func intSliceConv(s string) []int {
	res := make([]int, 0)
	vals := strings.Split(s, "")
	for _, v := range vals {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func Run(filename string) (Result, error) {
	batteries := utils.OpenFileIntoSlice(filename, intSliceConv)
	return Result{
		PartOne: partOne(batteries),
		PartTwo: partTwo(batteries),
	}, nil
}

func partOne(batteries [][]int) int {
	sum := 0
	for _, battery := range batteries {
		i := 0
		j := 1
		for idx := 1; idx < len(battery); idx++ {
			switch {
			// later entry is greater than first
			case battery[idx] > battery[i] && idx != len(battery)-1:
				i = idx
				j = idx + 1
				// later entry is greater than second but less than or equal to first
			case battery[idx] > battery[j]:
				j = idx
			}
		}
		v := (10 * battery[i]) + battery[j]
		sum += v
	}
	return sum
}

func partTwo(batteries [][]int) int {
	sum := 0
	for _, battery := range batteries {
		num := make([]int, 12)
		startIdx := 0
		for idx := range num {
			endIdx := len(battery) - (12 - idx - 1)
			nextNum, offset := maxWithIdx(battery[startIdx:endIdx])
			startIdx += offset
			num[idx] = nextNum
		}
		sum += convToNumber(num)
	}
	return sum
}

func maxWithIdx(digits []int) (int, int) {
	max, idx := -1, -1
	for i, v := range digits {
		if v > max {
			max = v
			idx = i
		}
	}
	return max, idx + 1
}

func convToNumber(d []int) int {
	b := make([]string, 0, len(d))
	for _, v := range d {
		b = append(b, strconv.Itoa(v))
	}
	i, err := strconv.Atoi(strings.Join(b, ""))
	if err != nil {
		panic(fmt.Sprintf("not a number: %s", b))
	}
	return i
}
