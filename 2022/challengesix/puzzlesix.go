package challengesix

import (
	"AdventOfCode/2022/file"
	"fmt"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
	ins := file.OpenFileIntoStringSlice(filename, file.StringConv)
	input := ins[0]
	res := Result{
		PartOne: idxOfRun(input, 4),
		PartTwo: idxOfRun(input, 14),
	}
	return res, nil
}

func idxOfRun(input string, subStrLen int) int {
	vals := make(map[rune]int, 26)
	for idx, ch := range input {
		if idx < subStrLen-1 {
			vals[ch] += 1
			continue
		}
		vals[ch] += 1
		if len(vals) == subStrLen {
			return idx + 1
		}
		vals[rune(input[idx-(subStrLen-1)])] -= 1
		if vals[rune(input[idx-(subStrLen-1)])] == 0 {
			delete(vals, rune(input[idx-(subStrLen-1)]))
		}

	}
	return -1
}

// changed this function to the one above after discussion on gophers slack
func idxOfRunOriginal(input string, subStrLen int) int {
	vals := make(map[rune]int, 26)
	for idx, ch := range input {
		if idx < subStrLen-1 {
			vals[ch] += 1
			continue
		}
		vals[ch] += 1
		if all(vals) {
			return idx + 1
		}
		vals[rune(input[idx-(subStrLen-1)])] -= 1

	}
	return -1
}

func all(m map[rune]int) bool {
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
