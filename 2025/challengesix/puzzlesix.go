package challengesix

import (
	"AdventOfCode/utils"
	"fmt"
	"slices"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
	return Result{
		PartOne: partOne(filename),
		PartTwo: partTwo(filename),
	}, nil
}

var opMap = map[string]func(int, int) int{
	"*": multiply,
	"+": add,
}

func partOne(filename string) int {
	rows := utils.OpenFileIntoSlice(filename, utils.StringRowConv(" "))
	for idx := range rows {
		rows[idx] = slices.DeleteFunc(rows[idx], func(s string) bool {
			return s == ""
		})
	}
	nums := make([][]int, 0, len(rows)-1)
	for idx := range rows {
		if idx == len(rows)-1 {
			continue
		}
		nums = append(nums, utils.Map(rows[idx], func(s string) int {
			return utils.IntConv(s)
		}))
	}
	ops := rows[len(rows)-1]
	sum := 0
	for i := range nums[0] {
		opFn := opMap[ops[i]]
		v := nums[0][i]
		for j := 1; j < len(nums); j++ {
			v = opFn(v, nums[j][i])
		}
		sum += v
	}
	return sum
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func partTwo(filename string) int {
	rows := utils.OpenFileIntoSlice(filename, utils.StringRowConv(""))
	sum := 0

	opRow := len(rows) - 1
	var opFn func(int, int) int
	curNums := make([]int, 0, len(rows)-1)

	for col := 0; col < len(rows[0]); col++ {
		if rows[opRow][col] != " " {
			if len(curNums) > 0 {
				v := curNums[0]
				for i := 1; i < len(curNums); i++ {
					v = opFn(v, curNums[i])
				}
				sum += v
			}

			opFn = opMap[rows[opRow][col]]
			curNums = curNums[:0]
		}
		curNum := ""
		for row := 0; row < opRow; row++ {
			if rows[row][col] != " " {
				curNum += string(rows[row][col])
			}
		}

		if curNum != "" {
			curNums = append(curNums, utils.IntConv(curNum))
		}
	}

	// handle the last set of numbers
	v := curNums[0]
	for i := 1; i < len(curNums); i++ {
		v = opFn(v, curNums[i])
	}
	sum += v
	return sum
}
