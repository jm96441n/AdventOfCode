package challengefour

import (
	"AdventOfCode/utils"
	"fmt"
	"iter"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
	warehouse := utils.OpenFileIntoSlice(filename, utils.StringRowConv(""))
	return Result{
		PartOne: partOne(warehouse),
		PartTwo: partTwo(warehouse),
	}, nil
}

func partOne(warehouse [][]string) int {
	sum := 0

	maxRow := len(warehouse)
	maxCol := len(warehouse[0])
	for i := 0; i < ((len(warehouse)) * (len(warehouse[0]))); i++ {
		row, col := utils.DivMod(i, len(warehouse[0]))
		if warehouse[row][col] != "@" {
			continue
		}

		nearbyRolls := 0
		for nb := range neighbors(row, col, maxRow, maxCol) {
			if warehouse[nb[0]][nb[1]] == "@" {
				nearbyRolls++
			}
		}
		if nearbyRolls < 4 {
			sum++
		}
	}
	return sum
}

func partTwo(warehouse [][]string) int {
	sum := 0

	maxRow := len(warehouse)
	maxCol := len(warehouse[0])
	for {
		seen := make(map[[2]int]struct{})
		for i := 0; i < ((len(warehouse)) * (len(warehouse[0]))); i++ {
			row, col := utils.DivMod(i, len(warehouse[0]))
			if warehouse[row][col] != "@" {
				continue
			}

			nearbyRolls := 0
			for nb := range neighbors(row, col, maxRow, maxCol) {
				if warehouse[nb[0]][nb[1]] == "@" {
					nearbyRolls++
				}
			}
			if nearbyRolls < 4 {
				seen[[2]int{row, col}] = struct{}{}
				sum++
			}
		}
		for p := range seen {
			warehouse[p[0]][p[1]] = "x"
		}
		if len(seen) == 0 {
			break
		}
	}
	return sum
}

var (
	rowDelta = []int{-1, -1, -1, 0, 1, 1, 1, 0}
	colDelta = []int{-1, 0, 1, 1, 1, 0, -1, -1}
)

func neighbors(row, col, maxRow, maxCol int) iter.Seq[[2]int] {
	return func(yield func([2]int) bool) {
		for i := range rowDelta {
			newRow := row + rowDelta[i]
			newCol := col + colDelta[i]
			if newRow >= 0 && newRow < maxRow && newCol >= 0 && newCol < maxCol {
				if !yield([2]int{newRow, newCol}) {
					return
				}
			}

		}
	}
}
