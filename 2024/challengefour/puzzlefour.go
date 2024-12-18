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
	return Result{
		PartOne: p1,
	}, nil
}

func neighbors(row, col, maxRow, maxCol int) [][2]int {
	rowDelta := []int{-1, 0, 1, 0, -1, 1, -1, 1}
	colDelta := []int{0, -1, 0, 1, -1, 1, 1, -1}
	nbrs := make([][2]int, 0, 4)
	for i := 0; i < len(rowDelta); i++ {
		newRow := col + rowDelta[i]
		newCol := row + colDelta[i]
		if newRow >= 0 && newRow < maxRow && newCol >= 0 && newCol < maxCol {
			nbrs = append(nbrs, [2]int{newRow, newCol})
		}
	}
	return nbrs
}

var vals = []byte{'X', 'M', 'A', 'S'}

func dfs(row, col, depth int, input []string) int {
	if vals[depth] != input[row][col] {
		return 0
	}

	if depth == 3 {
		return 1
	}

	curSum := 0
	for _, pair := range neighbors(row, col, len(input), len(input[0])) {
		curSum += dfs(pair[0], pair[1], depth+1, input)
	}
	return curSum
}

func partOne(input []string) int {
	sum := 0
	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == 'X' {
				sum += dfs(row, col, 0, input)
			}
		}
	}
	return sum
}
