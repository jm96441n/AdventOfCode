package challenge11

import (
	"AdventOfCode/utils"
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

type changedPosition struct {
	position
	newVal string
}

type adjacentPositionFn func(pos position) position
type occupiedFn func(posFunc adjacentPositionFn, curPos position, rows [][]string) bool

func Run() {
	rows := utils.OpenFileIntoStringSlice("challenge11/input.txt")
	splitRows := make([][]string, len(rows))
	for idx, row := range rows {
		splitRows[idx] = strings.Split(row, "")
	}
	ptone := make(chan int, 1)
	pttwo := make(chan int, 1)
	go func() {
		unOccupied := conwaysGameOfSeats(splitRows, isAdjacentOccupied, 4)
		ptone <- unOccupied
	}()
	go func() {
		unOccupied := conwaysGameOfSeats(splitRows, isFirstSeatInRowOccupied, 5)
		pttwo <- unOccupied
	}()

	ptOneUn := <-ptone
	ptTwoUn := <-pttwo
	fmt.Println(ptOneUn)
	fmt.Println(ptTwoUn)
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

func conwaysGameOfSeats(inputs [][]string, occupiedChecker occupiedFn, countToEmpty int) int {
	rows := runTheGame(inputs, occupiedChecker, countToEmpty)
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

func runTheGame(inputs [][]string, occupiedChecker occupiedFn, countToEmpty int) [][]string {
	rows := make([][]string, len(inputs))
	for idx, row := range inputs {
		rows[idx] = append(make([]string, 0, len(row)), row...)
	}

	for {
		changes := make([]changedPosition, 0)
		for rowIdx, row := range rows {
			for colIdx, seat := range row {
				curPos := position{
					row: rowIdx,
					col: colIdx,
				}
				occupiedCount := 0
				for _, fn := range adjPosFns {
					if occupiedChecker(fn, curPos, rows) {
						occupiedCount++
					}
				}
				if seat == empty && occupiedCount == 0 {
					changes = append(changes, changedPosition{position: curPos, newVal: taken})
				}

				if seat == taken && occupiedCount >= countToEmpty {
					changes = append(changes, changedPosition{position: curPos, newVal: empty})
				}

			}
		}

		if len(changes) == 0 {
			break
		}
		for _, pos := range changes {
			rows[pos.row][pos.col] = pos.newVal
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

func isFirstSeatInRowOccupied(posFunc adjacentPositionFn, curPos position, rows [][]string) bool {
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

	return isFirstSeatInRowOccupied(posFunc, nextPos, rows)
}

func isAdjacentOccupied(posFunc adjacentPositionFn, curPos position, rows [][]string) bool {
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
