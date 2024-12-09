package challenge17

import (
	"AdventOfCode/utils"
	"fmt"
	"strings"
)

var (
	x        = "x"
	y        = "y"
	z        = "z"
	inactive = "."
	active   = "#"
)

func Run() {
	rows := utils.OpenFileIntoStringSlice("challenge17/input.txt")
	fmt.Println(run3DBootCycle(rows))
	fmt.Println(run4DBootCycle(rows))
}

func run3DBootCycle(rows []string) int {
	matrix := make([][][]string, len(rows))
	for xidx, row := range rows {
		splitRow := strings.Split(row, "")
		matrix[xidx] = make([][]string, len(row))
		for yidx, char := range splitRow {
			matrix[xidx][yidx] = make([]string, 1)
			matrix[xidx][yidx][0] = char
		}
	}
	matrix = grow3D(matrix)
	for i := 0; i < 6; i++ {
		changeSpots := make([][]int, 0)
		for xidx := range matrix {
			for yidx := range matrix[xidx] {
				for zidx := range matrix[xidx][yidx] {
					nc := active3DNeighborCount(xidx, yidx, zidx, matrix)
					if matrix[xidx][yidx][zidx] == active {
						if nc < 2 || nc > 3 {
							changeSpots = append(changeSpots, []int{xidx, yidx, zidx})
						}
					} else {
						if nc == 3 {
							changeSpots = append(changeSpots, []int{xidx, yidx, zidx})
						}
					}
				}
			}
		}
		for _, changeTuple := range changeSpots {
			if matrix[changeTuple[0]][changeTuple[1]][changeTuple[2]] == active {
				matrix[changeTuple[0]][changeTuple[1]][changeTuple[2]] = inactive
			} else {
				matrix[changeTuple[0]][changeTuple[1]][changeTuple[2]] = active
			}
		}
		matrix = grow3D(matrix)
	}
	activeCount := 0
	for xidx := range matrix {
		for yidx := range matrix[xidx] {
			for zidx := range matrix[xidx][yidx] {
				if matrix[xidx][yidx][zidx] == active {
					activeCount++
				}
			}
		}
	}
	return activeCount
}

func active3DNeighborCount(curX, curY, curZ int, matrix [][][]string) int {
	count := 0
	for xc := -1; xc <= 1; xc++ {
		for yc := -1; yc <= 1; yc++ {
			for zc := -1; zc <= 1; zc++ {
				if xc == 0 && yc == 0 && zc == 0 {
					continue
				}
				if curX-xc < 0 || curY-yc < 0 || curZ-zc < 0 {
					continue
				}
				if curX-xc >= len(matrix) || curY-yc >= len(matrix[0]) || curZ-zc >= len(matrix[0][0]) {
					continue
				}
				if matrix[curX-xc][curY-yc][curZ-zc] == active {
					count++
				}
			}
		}
	}
	return count
}

func grow3D(prevArr [][][]string) [][][]string {
	newArr := make([][][]string, len(prevArr)+4)
	for xidx := range newArr {
		outsideXRange := (xidx < 2) || ((xidx - 2) >= len(prevArr))
		newArr[xidx] = make([][]string, len(prevArr[0])+4)
		for yidx := range newArr[xidx] {
			outsideYRange := (yidx < 2) || ((yidx - 2) >= len(prevArr[0]))
			newArr[xidx][yidx] = make([]string, len(prevArr[0][0])+4)
			for zidx := range newArr[xidx][yidx] {
				outsideZRange := (zidx < 2) || ((zidx - 2) >= len(prevArr[0][0]))
				if outsideXRange || outsideYRange || outsideZRange {
					newArr[xidx][yidx][zidx] = inactive
				} else {
					newArr[xidx][yidx][zidx] = prevArr[xidx-2][yidx-2][zidx-2]
				}
			}
		}
	}
	return newArr
}

func run4DBootCycle(rows []string) int {
	matrix := make([][][][]string, len(rows))
	for xidx, row := range rows {
		splitRow := strings.Split(row, "")
		matrix[xidx] = make([][][]string, len(row))
		for yidx, char := range splitRow {
			matrix[xidx][yidx] = make([][]string, 1)
			matrix[xidx][yidx][0] = make([]string, 1)
			matrix[xidx][yidx][0][0] = char
		}
	}
	matrix = grow4D(matrix)
	for i := 0; i < 6; i++ {
		changeSpots := make([][]int, 0)
		for xidx := range matrix {
			for yidx := range matrix[xidx] {
				for zidx := range matrix[xidx][yidx] {
					for widx := range matrix[xidx][yidx][zidx] {
						nc := active4DNeighborCount(xidx, yidx, zidx, widx, matrix)
						if matrix[xidx][yidx][zidx][widx] == active {
							if nc < 2 || nc > 3 {
								changeSpots = append(changeSpots, []int{xidx, yidx, zidx, widx})
							}
						} else {
							if nc == 3 {
								changeSpots = append(changeSpots, []int{xidx, yidx, zidx, widx})
							}
						}
					}
				}
			}
		}
		for _, changeTuple := range changeSpots {
			if matrix[changeTuple[0]][changeTuple[1]][changeTuple[2]][changeTuple[3]] == active {
				matrix[changeTuple[0]][changeTuple[1]][changeTuple[2]][changeTuple[3]] = inactive
			} else {
				matrix[changeTuple[0]][changeTuple[1]][changeTuple[2]][changeTuple[3]] = active
			}
		}
		matrix = grow4D(matrix)
	}
	activeCount := 0
	for xidx := range matrix {
		for yidx := range matrix[xidx] {
			for zidx := range matrix[xidx][yidx] {
				for widx := range matrix[xidx][yidx][zidx] {
					if matrix[xidx][yidx][zidx][widx] == active {
						activeCount++
					}

				}
			}
		}
	}
	return activeCount
}

func active4DNeighborCount(curX, curY, curZ, curW int, matrix [][][][]string) int {
	count := 0
	for xc := -1; xc <= 1; xc++ {
		for yc := -1; yc <= 1; yc++ {
			for zc := -1; zc <= 1; zc++ {
				for wc := -1; wc <= 1; wc++ {
					if xc == 0 && yc == 0 && zc == 0 && wc == 0 {
						continue
					}
					if curX-xc < 0 || curY-yc < 0 || curZ-zc < 0 || curW-wc < 0 {
						continue
					}
					if curX-xc >= len(matrix) || curY-yc >= len(matrix[0]) || curZ-zc >= len(matrix[0][0]) || curW-wc >= len(matrix[0][0][0]) {
						continue
					}
					if matrix[curX-xc][curY-yc][curZ-zc][curW-wc] == active {
						count++
					}
				}
			}
		}
	}
	return count
}

func grow4D(prevArr [][][][]string) [][][][]string {
	newArr := make([][][][]string, len(prevArr)+4)
	for xidx := range newArr {
		outsideXRange := (xidx < 2) || ((xidx - 2) >= len(prevArr))
		newArr[xidx] = make([][][]string, len(prevArr[0])+4)
		for yidx := range newArr[xidx] {
			outsideYRange := (yidx < 2) || ((yidx - 2) >= len(prevArr[0]))
			newArr[xidx][yidx] = make([][]string, len(prevArr[0][0])+4)
			for zidx := range newArr[xidx][yidx] {
				outsideZRange := (zidx < 2) || ((zidx - 2) >= len(prevArr[0][0]))
				newArr[xidx][yidx][zidx] = make([]string, len(prevArr[0][0][0])+4)
				for widx := range newArr[xidx][yidx][zidx] {
					outsideWRange := (widx < 2) || ((widx - 2) >= len(prevArr[0][0][0]))
					if outsideXRange || outsideYRange || outsideZRange || outsideWRange {
						newArr[xidx][yidx][zidx][widx] = inactive
					} else {
						newArr[xidx][yidx][zidx][widx] = prevArr[xidx-2][yidx-2][zidx-2][widx-2]
					}
				}
			}
		}
	}
	return newArr
}
