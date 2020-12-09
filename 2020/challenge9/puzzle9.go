package challenge9

import (
	"AdventOfCode/challenge1"
	"AdventOfCode/file_utils"
	"fmt"
	"strconv"
)

var preambleLen int = 5

func Run() {
	rows := file_utils.OpenFileIntoSlice("test_input.txt")
	intRows := make([]int, len(rows))
	for idx, num := range rows {
		intRow, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		intRows[idx] = intRow
	}
	noMatch := 0
	for idx, row := range intRows[preambleLen:] {
		_, addendTwo := challenge1.TwoSum(intRows[(idx-preambleLen):idx], row)
		if addendTwo == 0 {
			noMatch = row
			break
		}
	}
	fmt.Println(noMatch)

}
