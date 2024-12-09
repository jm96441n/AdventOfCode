package challenge10

import (
	"AdventOfCode/utils"
	"fmt"
	"sort"
)

func Run() {
	rows := utils.OpenFileIntoIntSlice("challenge10/input.txt")
	sort.Ints(rows)
	fmt.Println(productOfOnesAndThrees(rows))
	fmt.Println(combinations(rows))
}

func productOfOnesAndThrees(rows []int) int {
	oneDiff := 0
	threeDiff := 1
	prevRow := 0
	for _, row := range rows {
		diff := row - prevRow
		switch diff {
		case 1:
			oneDiff++
		case 3:
			threeDiff++
		default:
			if diff != 2 {
				panic("DIFF IS MORE THAN 3")
			}

		}
		prevRow = row
	}
	return (oneDiff * threeDiff)
}

func combinations(rows []int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	for _, row := range rows {
		for i := 1; i < 4; i++ {
			if val, ok := sumMap[(row - i)]; ok {
				sumMap[row] += val
			} else {
				sumMap[row] += 0
			}
		}
	}
	return sumMap[rows[(len(rows)-1)]]
}
