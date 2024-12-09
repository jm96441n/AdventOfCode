package challengefive

import (
	"AdventOfCode/2022/file"
	"fmt"
)

type Result struct {
	PartOne string
	PartTwo string
}

func (r Result) Display() {
	fmt.Printf("PartOne: %s\nPartTwo: %s\n", r.PartOne, r.PartTwo)
}

type Move struct {
	Amount int
	Start  int
	Dest   int
}

func Run(filename string) (Result, error) {
	ins := file.OpenFileIntoStringSlice(filename, file.StringConv)
	numStacks := ((len(ins[0]) + 1) / 4)
	stacksOne := make([]string, numStacks)
	stacksTwo := make([]string, numStacks)
	var idx int
	for idx = 0; idx < len(ins); idx++ {
		line := ins[idx]
		if line[0:2] == " 1" {
			break
		}
		line_idx := 1
		stack_idx := 0
		for line_idx < len(line) {
			if line[line_idx] != ' ' {
				stacksOne[stack_idx] = string(line[line_idx]) + stacksOne[stack_idx]
				stacksTwo[stack_idx] = string(line[line_idx]) + stacksTwo[stack_idx]
			}
			stack_idx += 1
			line_idx += 4
		}
	}
	idx += 2
	moves := make([]Move, 0)
	for idx < len(ins) {
		move := Move{}
		fmt.Sscanf(ins[idx], "move %d from %d to %d", &move.Amount, &move.Start, &move.Dest)
		moves = append(moves, move)
		idx++
	}

	for _, move := range moves {
		stackToMoveIdx := move.Start - 1
		stacksOne[move.Dest-1] += reverse(stacksOne[stackToMoveIdx][len(stacksOne[stackToMoveIdx])-move.Amount : len(stacksOne[stackToMoveIdx])])
		stacksOne[stackToMoveIdx] = stacksOne[stackToMoveIdx][0 : len(stacksOne[stackToMoveIdx])-move.Amount]
		stacksTwo[move.Dest-1] += stacksTwo[stackToMoveIdx][len(stacksTwo[stackToMoveIdx])-move.Amount : len(stacksTwo[stackToMoveIdx])]
		stacksTwo[stackToMoveIdx] = stacksTwo[stackToMoveIdx][0 : len(stacksTwo[stackToMoveIdx])-move.Amount]

	}
	pOne := ""
	pTwo := ""
	for i := 0; i < len(stacksOne); i++ {
		sOne := stacksOne[i]
		sTwo := stacksTwo[i]
		if len(sOne) > 0 {
			pOne += string(sOne[len(sOne)-1])
		}
		if len(sTwo) > 0 {
			pTwo += string(sTwo[len(sTwo)-1])
		}

	}

	return Result{PartOne: pOne, PartTwo: pTwo}, nil
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// N = (11 + 1) / 4
