package challengeseven

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
	rows := utils.OpenFileIntoSlice(filename, utils.StringRowConv(""))
	return Result{
		PartOne: partOne(rows),
		PartTwo: partTwo(rows),
	}, nil
}

type position struct {
	row int
	col int
}

func partOne(rows [][]string) int {
	q := utils.NewQueue[position]()
	for i := range len(rows[0]) {
		if rows[0][i] == "S" {
			q.Enqueue(position{0, i})
		}
	}

	splits := make(map[position]struct{})
	seen := make(map[position]struct{})
	for !q.IsEmpty() {
		p, _ := q.Dequeue()
		if rows[p.row][p.col] == "^" {
			splits[p] = struct{}{}
			left, right := position{p.row, p.col - 1}, position{p.row, p.col + 1}

			if _, ok := seen[left]; !ok {
				q.Enqueue(left)
			}

			if _, ok := seen[right]; !ok {
				q.Enqueue(right)
			}

		} else if p.row+1 < len(rows) {
			if _, ok := seen[position{p.row + 1, p.col}]; !ok {
				q.Enqueue(position{p.row + 1, p.col})
			}
		}

		seen[p] = struct{}{}
	}

	return len(splits)
}

func partTwo(rows [][]string) int {
	// DFS this time around
	return 0
}
