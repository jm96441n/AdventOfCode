package challenge15

import (
	"fmt"
	"strconv"
	"strings"
)

var testInput string = "0,3,6"
var puzzleInput string = "15,5,1,4,7,0"

var partOneGoal int = 2020
var partTwoGoal int = 30000000

func Run() {
	numMap := make(map[int]int)
	rows := strings.Split(puzzleInput, ",")
	var prevNum int
	for idx, snum := range rows[:(len(rows) - 1)] {
		num, err := strconv.Atoi(snum)
		if err != nil {
			panic(err)
		}

		numMap[num] = (idx + 1)
	}
	prevNum, err := strconv.Atoi(rows[(len(rows) - 1)])
	if err != nil {
		panic(err)
	}

	fmt.Println(findXthNumberSpokenIt(numMap, prevNum, len(rows)+1, partTwoGoal))
}

func findXthNumberSpokenIt(numMap map[int]int, prevNum, startTurn, maxTurns int) int {
	var spoken int
	for curTurn, prevTurn := startTurn, startTurn-1; curTurn <= maxTurns; curTurn, prevTurn = curTurn+1, prevTurn+1 {
		spoken = 0
		if turn, ok := numMap[prevNum]; ok {
			spoken = prevTurn - turn
		}
		numMap[prevNum] = prevTurn
		prevNum = spoken
	}
	return spoken

}

func findXthNumberSpokenRec(numMap map[int]int, prevNum, curTurn, maxTurns int) int {
	if curTurn == maxTurns {
		return prevNum
	}

	var curNum int

	if turn, ok := numMap[prevNum]; ok {
		curNum = curTurn - turn
	}

	numMap[prevNum] = curTurn
	curTurn++
	return findXthNumberSpokenRec(numMap, curNum, curTurn, maxTurns)
}
