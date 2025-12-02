package challengeone

import (
	"fmt"
	"math"
	"strconv"

	"AdventOfCode/utils"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

type rotation struct {
	direction string
	distance  int
}

func rotationConv(s string) rotation {
	direction := string(s[0])
	distance, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}
	return rotation{direction: direction, distance: distance}
}

func Run(filename string) (Result, error) {
	rotations := utils.OpenFileIntoSlice(filename, rotationConv)
	partOnePW := partOne(rotations)
	partTwoPw := partTwo(rotations)
	return Result{PartOne: partOnePW, PartTwo: partTwoPw}, nil
}

func partOne(rotations []rotation) int {
	pos := 50
	pw := 0

	for _, rot := range rotations {
		dist := rot.distance
		if rot.direction == "L" {
			dist = -dist
		}
		pos += dist

		pos := mod(pos, 100)

		if pos == 0 {
			pw += 1
		}
	}
	return pw
}

func partTwo(rotations []rotation) int {
	pos := 50
	pw := 0

	for _, rot := range rotations {
		dist := rot.distance
		if rot.direction == "L" {
			dist = -dist
		}
		newPos := pos + dist
		if dist >= 0 {
			pw += divFloor(newPos, 100) - divFloor(pos, 100)
		} else {
			pw += divFloor(pos-1, 100) - divFloor(newPos-1, 100)
		}
		pos = mod(newPos, 100)

	}
	return pw
}

func mod(a, b int) int {
	d := divFloor(a, b)
	return a - (b * d)
}

func divFloor(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}
