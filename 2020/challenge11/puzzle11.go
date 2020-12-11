package challenge11

import (
	"AdventOfCode/file_utils"
	"fmt"
	"strings"
)

var (
	empty = "L"
	taken = "#"
	floor = "."
)

type position struct {
	row int
	col int
}

type adjacentPositionFn func(pos position) position
type gameExecutor func(inputs [][]string, occupiedChecker occupiedFn) [][]string
type occupiedFn func(posFunc adjacentPositionFn, curPos position, rows [][]string) bool

func Run() {
	rows := file_utils.OpenFileIntoSlice("challenge11/input.txt")
	splitRows := make([][]string, 0)
	for _, row := range rows {
		splitRows = append(splitRows, strings.Split(row, ""))
	}
	unOccupied := conwaysGameOfSeats(splitRows, partOneSeats, isOccupied)
	fmt.Println(unOccupied)
	ptTwoUnoccupied := conwaysGameOfSeats(splitRows, partTwoSeats, isNextOccupied)
	fmt.Println(ptTwoUnoccupied)
}

var adjPosFns = []adjacentPositionFn{
	topPosition,
	topRightPosition,
	rightPosition,
	downRightPosition,
	downPosition,
	downLeftPosition,
	leftPosition,
	topLeftPosition,
}

func conwaysGameOfSeats(inputs [][]string, gameFn gameExecutor, occupiedChecker occupiedFn) int {
	rows := gameFn(inputs, occupiedChecker)
	occupiedSeatCount := 0
	for _, row := range rows {
		for _, seat := range row {
			if seat == taken {
				occupiedSeatCount++
			}
		}

	}
	return occupiedSeatCount
}

func partOneSeats(inputs [][]string, occupiedChecker occupiedFn) [][]string {
	rows := make([][]string, len(inputs))
	for idx, row := range inputs {
		rows[idx] = append(make([]string, 0, len(row)), row...)
	}

	prevArrangment := make([][]string, len(rows))
	for {
		for idx, row := range rows {
			prevArrangment[idx] = append(make([]string, 0, len(row)), row...)
		}
		for rowIdx, row := range prevArrangment {
			for colIdx, seat := range row {
				curPos := position{
					row: rowIdx,
					col: colIdx,
				}
				occupiedCount := 0
				for _, fn := range adjPosFns {
					if occupiedChecker(fn, curPos, prevArrangment) {
						occupiedCount++
					}
				}
				if seat == empty && occupiedCount == 0 {
					rows[rowIdx][colIdx] = taken
				}
				if seat == taken && occupiedCount >= 4 {
					rows[rowIdx][colIdx] = empty
				}

			}
		}
		matching := slicesDeepEqual(rows, prevArrangment)
		if matching {
			break
		}

	}
	return rows
}

func partTwoSeats(inputs [][]string, occupiedChecker occupiedFn) [][]string {
	rows := make([][]string, len(inputs))
	for idx, row := range inputs {
		rows[idx] = append(make([]string, 0, len(row)), row...)
	}

	prevArrangment := make([][]string, len(rows))
	for {
		for idx, row := range rows {
			prevArrangment[idx] = append(make([]string, 0, len(row)), row...)
		}
		for rowIdx, row := range prevArrangment {
			for colIdx, seat := range row {
				curPos := position{
					row: rowIdx,
					col: colIdx,
				}
				occupiedCount := 0
				for _, fn := range adjPosFns {
					if occupiedChecker(fn, curPos, prevArrangment) {
						occupiedCount++
					}
				}
				if seat == empty && occupiedCount == 0 {
					rows[rowIdx][colIdx] = taken
				}
				if seat == taken && occupiedCount >= 5 {
					rows[rowIdx][colIdx] = empty
				}

			}
		}
		matching := slicesDeepEqual(rows, prevArrangment)
		if matching {
			break
		}

	}
	return rows
}

func slicesDeepEqual(rows, prevArrangment [][]string) bool {
	matching := true
	for idx, prevRow := range prevArrangment {
		if strings.Join(prevRow, "") != strings.Join(rows[idx], "") {
			matching = false
			break
		}
	}
	return matching
}

func isNextOccupied(posFunc adjacentPositionFn, curPos position, rows [][]string) bool {
	nextPos := posFunc(curPos)
	if posOutsideRows(len(rows), len(rows[0]), nextPos) {
		return false
	}
	if rows[nextPos.row][nextPos.col] == taken {
		return true
	}

	if rows[nextPos.row][nextPos.col] == empty {
		return false
	}

	return isNextOccupied(posFunc, nextPos, rows)
}

func isOccupied(posFunc adjacentPositionFn, curPos position, rows [][]string) bool {
	adjPos := posFunc(curPos)
	if posOutsideRows(len(rows), len(rows[0]), adjPos) {
		return false
	}
	if rows[adjPos.row][adjPos.col] == taken {
		return true
	}
	return false

}

func posOutsideRows(maxRowIdx, maxColIdx int, adjPos position) bool {
	if adjPos.col < 0 || adjPos.row < 0 {
		return true
	}

	if adjPos.col >= maxColIdx || adjPos.row >= maxRowIdx {
		return true
	}
	return false
}

func topPosition(pos position) position {
	return position{
		row: pos.row - 1,
		col: pos.col,
	}
}

func topRightPosition(pos position) position {
	return position{
		row: pos.row - 1,
		col: pos.col + 1,
	}
}

func rightPosition(pos position) position {
	return position{
		row: pos.row,
		col: pos.col + 1,
	}
}

func downRightPosition(pos position) position {
	return position{
		row: pos.row + 1,
		col: pos.col + 1,
	}
}

func downPosition(pos position) position {
	return position{
		row: pos.row + 1,
		col: pos.col,
	}
}

func downLeftPosition(pos position) position {
	return position{
		row: pos.row + 1,
		col: pos.col - 1,
	}
}

func leftPosition(pos position) position {
	return position{
		row: pos.row,
		col: pos.col - 1,
	}
}

func topLeftPosition(pos position) position {
	return position{
		row: pos.row - 1,
		col: pos.col - 1,
	}
}
