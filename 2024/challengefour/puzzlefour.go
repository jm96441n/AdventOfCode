package challengefour

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
	input := utils.OpenFileIntoSlice(filename, utils.StringConv)
	p1 := partOne(input)
	p2 := partTwo(input)
	return Result{
		PartOne: p1,
		PartTwo: p2,
	}, nil
}

var vals = []byte{'X', 'M', 'A', 'S'}

func partOne(input []string) int {
	rowDelta := []int{-1, 0, 1, 0, -1, 1, -1, 1}
	colDelta := []int{0, -1, 0, 1, -1, 1, 1, -1}
	sum := 0
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == 'X' {
				for i := 0; i < len(rowDelta); i++ {
					curRow := row
					curCol := col
					found := true
					for j := 0; j < len(vals); j++ {
						if (curRow < 0 || curRow >= len(input) || curCol < 0 || curCol >= len(input[0])) || input[curRow][curCol] != vals[j] {
							found = false
							break
						}
						curRow += rowDelta[i]
						curCol += colDelta[i]
					}
					if found {
						sum += 1
					}
				}
			}
		}
	}
	return sum
}

func partTwo(input []string) int {
	sum := 0

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == 'A' {
				uld := []int{row - 1, col - 1}
				lrd := []int{row + 1, col + 1}

				urd := []int{row - 1, col + 1}
				lld := []int{row + 1, col - 1}

				valid := true
				for _, pair := range [][]int{uld, lrd, urd, lld} {
					if pair[0] < 0 || pair[0] >= len(input) || pair[1] < 0 || pair[1] >= len(input[0]) {
						valid = false
						break
					}
				}
				if !valid {
					continue
				}

				rightDiag := fmt.Sprintf("%sA%s", string(input[uld[0]][uld[1]]), string(input[lrd[0]][lrd[1]]))
				leftDiag := fmt.Sprintf("%sA%s", string(input[lld[0]][lld[1]]), string(input[urd[0]][urd[1]]))
				if (rightDiag == "MAS" || rightDiag == "SAM") && (leftDiag == "MAS" || leftDiag == "SAM") {
					sum += 1
				}
			}
		}
	}

	return sum
}
