package challengeone

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"os"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(fileName string) (Result, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return Result{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	col1 := make([]int, 0)
	col2 := make([]int, 0)
	for scanner.Scan() {
		a, b := 0, 0
		fmt.Sscanf(scanner.Text(), "%d  %d", &a, &b)
		col1 = append(col1, a)
		col2 = append(col2, b)
	}

	if err := scanner.Err(); err != nil {
		return Result{}, err
	}

	return Result{
		PartOne: PartOne(col1, col2),
		PartTwo: PartTwo(col1, col2),
	}, nil
}

func PartOne(colOne, colTwo []int) int {
	colOnePQ := utils.ToPriorityQueue(colOne)
	colTwoPQ := utils.ToPriorityQueue(colTwo)
	sum := 0
	for colOnePQ.Size > 0 {
		v1 := colOnePQ.Pop()
		v2 := colTwoPQ.Pop()
		sum += utils.AbsValue(v1 - v2)
	}
	return sum
}

func PartTwo(colOne, colTwo []int) int {
	counts := make(map[int]int)
	for _, n := range colTwo {
		counts[n] += 1
	}

	sum := 0
	for _, n := range colOne {
		sum += (n * counts[n])
	}
	return sum
}
