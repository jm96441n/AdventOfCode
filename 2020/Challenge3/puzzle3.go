package challenge3

import (
	"AdventOfCode/file_utils"
	"fmt"
	"strings"
)

var tree string = "#"
var emptySpace string = "."
var visitedEmpty string = "0"
var visitedTree string = "X"

func Run() {
	rows := file_utils.OpenFileIntoSlice("./challenge3/input.txt")
	splitRows := make([][]string, len(rows))
	for idx, row := range rows {
		splitRows[idx] = strings.Split(row, "")
	}

	partOneSlope := []int{3, 1}
	slopes := [][]int{{1, 1}, {5, 1}, {7, 1}, {1, 2}}
	treeProduct := getTreesForSlope(partOneSlope[0], partOneSlope[1], splitRows)
	fmt.Println(treeProduct)
	for _, slope := range slopes {
		treeProduct *= getTreesForSlope(slope[0], slope[1], splitRows)
	}

	fmt.Println(treeProduct)
}

func getTreesForSlope(x, y int, splitRows [][]string) int {
	treeCount := 0
	for curHorPos, curVertPos := x, y; curVertPos < len(splitRows); curHorPos, curVertPos = curHorPos+x, curVertPos+y {
		horPos := curHorPos % len(splitRows[0])
		if splitRows[curVertPos][horPos] == tree || splitRows[curVertPos][horPos] == visitedTree {
			treeCount++
			splitRows[curVertPos][horPos] = visitedTree
		} else {
			splitRows[curVertPos][horPos] = visitedEmpty
		}
	}
	return treeCount
}
