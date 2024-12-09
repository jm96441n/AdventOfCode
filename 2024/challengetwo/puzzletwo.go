package challengetwo

import (
	"AdventOfCode/utils"
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
	rows := utils.OpenFileIntoSlice(filename, utils.IntRowConv)
	partOne := PartOne(rows)
	partTwo := PartTwo(rows)
	return Result{
		PartOne: partOne,
		PartTwo: partTwo,
	}, nil
}

func PartOne(rows [][]int) int {
	safeRows := 0
	for _, row := range rows {
		if isSafe(row) {
			safeRows++
		}
	}

	return safeRows
}

func PartTwo(rows [][]int) int {
	safeRows := 0
	for _, row := range rows {
		if isSafe(row) {
			safeRows++
			continue
		}
		for i := 0; i < len(row); i++ {
			b := make([]int, 0)
			b = append(b, row[0:i]...)
			if isSafe(append(b, row[i+1:]...)) {
				safeRows++
				break
			}
		}
	}

	return safeRows
}

func isSafe(row []int) bool {
	safe := true
	ascending := false
	descending := false
	for i := 1; i < len(row); i++ {
		if row[i-1] > row[i] {
			descending = true
		} else {
			ascending = true
		}

		// we were previously ascending and are now descending or vice versa
		if ascending == descending {
			safe = false
			break
		}

		diff := row[i-1] - row[i]

		if utils.AbsValue(diff) > 3 || diff == 0 {
			safe = false
			break
		}
	}
	return safe
}
