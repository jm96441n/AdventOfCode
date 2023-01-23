package challengeeight

import (
	"AdventOfCode/2022/file"
	"fmt"
	"sort"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
	treeRows := file.OpenFileIntoSlice(filename, file.StringConv)
	trees := make([][]int, len(treeRows))
	for idx, row := range treeRows {
		trees[idx] = make([]int, len(row))
		for colIdx, treeRune := range row {
			height := int(treeRune)
			trees[idx][colIdx] = height
		}
	}
	bfs := func(qu [][2]int, deltas [][2]int) map[[2]int]struct{} {
		reachable := make(map[[2]int]struct{}, 0)
		for len(qu) > 0 {
			pos := qu[0]
			qu = qu[1:]
			reachable[pos] = struct{}{}
			row, col := pos[0], pos[1]
			for _, delta := range deltas {
				nbrRow, nbrCol := row+delta[0], col+delta[1]
				if inGrid(trees, nbrRow, nbrCol) && mapHasKey(reachable, pos) && trees[nbrRow][nbrCol] > trees[row][col] {
					qu = append(qu, [2]int{nbrRow, nbrCol})
				}

			}

		}
		return reachable
	}
	leftRow := make([][2]int, len(trees))
	rightRow := make([][2]int, len(trees))
	topRow := make([][2]int, len(trees))
	bottomRow := make([][2]int, len(trees))

	for i := 0; i < len(trees); i++ {
		leftRow[i] = [2]int{i, 0}
		rightRow[i] = [2]int{i, len(trees[0]) - 1}
		topRow[i] = [2]int{0, i}
		bottomRow[i] = [2]int{len(trees) - 1, i}
	}
	leftVisible := bfs(leftRow, [][2]int{{1, 0}, {0, 1}})
	rightVisible := bfs(rightRow, [][2]int{{-1, 0}, {0, -1}})
	visible := make(map[[2]int]struct{}, 0)
	for k := range leftVisible {
		visible[k] = struct{}{}
	}
	for k := range rightVisible {
		visible[k] = struct{}{}
	}

	for i := 0; i < len(leftRow); i++ {
		visible[leftRow[i]] = struct{}{}
		visible[rightRow[i]] = struct{}{}
		visible[topRow[i]] = struct{}{}
		visible[bottomRow[i]] = struct{}{}
	}
	visibleSlice := make([][2]int, 0)
	for k := range visible {
		visibleSlice = append(visibleSlice, k)
	}
	sort.Slice(visibleSlice, func(x, y int) bool { return visibleSlice[x][0] < visibleSlice[y][0] })
	first := 0
	for _, k := range visibleSlice {
		if k[0] != first {
			fmt.Println()
			first = k[0]
		}
		fmt.Println(k)
	}

	return Result{PartOne: len(visible)}, nil
}

func inGrid(trees [][]int, nbrRow int, nbrCol int) bool {
	return 0 <= nbrRow && nbrRow < len(trees) && 0 <= nbrCol && nbrCol < len(trees[0])
}

func mapHasKey(m map[[2]int]struct{}, p [2]int) bool {
	_, found := m[p]
	return found
}
