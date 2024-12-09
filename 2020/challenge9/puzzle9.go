package challenge9

import (
	"AdventOfCode/2020/challenge1"
	"AdventOfCode/utils"
	"fmt"
	"strconv"
)

var preambleLen int = 25

func Run() {
	rows := utils.OpenFileIntoStringSlice("./challenge9/input.txt")
	intRows := make([]int, len(rows))
	for idx, num := range rows {
		intRow, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		intRows[idx] = intRow
	}
	noMatch := findNoMatch(intRows)
	sum := findSumNoMatchWindows(intRows, noMatch)

	fmt.Println(noMatch)
	fmt.Println(sum)
}

func findNoMatch(inputs []int) int {
	noMatch := 0
	for idx, row := range inputs {
		if idx < preambleLen {
			continue
		}
		_, addendTwo := challenge1.TwoSum(inputs[(idx-preambleLen):idx], row)
		if addendTwo == 0 {
			noMatch = row
			break
		}
	}
	return noMatch
}

func findSumNoMatchWindows(inputs []int, matcher int) int {
	highIdx := 0
	lowIdx := 0
	runSum := inputs[lowIdx]
	for {
		if runSum < matcher {
			highIdx++
			runSum += inputs[highIdx]
		} else if runSum > matcher {
			runSum -= inputs[lowIdx]
			lowIdx++
		} else {
			break
		}

	}
	fmt.Println(inputs[lowIdx:highIdx])
	potentialMatches := inputs[lowIdx:highIdx]
	low := potentialMatches[0]
	high := potentialMatches[0]
	for _, num := range potentialMatches {
		if num < low {
			low = num
		}

		if num > high {
			high = num
		}
	}
	return (low + high)
}

func findSumToNoMatch(inputs []int, matcher int) int {
	lowIdx := 0
	highIdx := 0
	for lowIdx = 0; lowIdx < len(inputs); lowIdx++ {
		runSum := inputs[lowIdx]
		for highIdx = lowIdx + 1; highIdx < len(inputs); highIdx++ {
			runSum += inputs[highIdx]
			if runSum >= matcher {
				break
			}
		}
		if runSum == matcher {
			break
		}
	}
	potentialMatches := inputs[lowIdx:highIdx]
	low := potentialMatches[0]
	high := potentialMatches[0]
	for _, num := range potentialMatches {
		if num < low {
			low = num
		}

		if num > high {
			high = num
		}
	}
	return (low + high)
}
