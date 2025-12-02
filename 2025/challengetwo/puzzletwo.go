package challengetwo

import (
	"fmt"
	"slices"
	"strings"

	"AdventOfCode/utils"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

type intRange struct {
	low  int
	high int
}

func Run(filename string) (Result, error) {
	idString := utils.OpenFileIntoString(filename)
	ids := strings.Split(idString, ",")
	intRanges := make([]intRange, 0, len(ids))
	for _, i := range ids {
		nums := strings.Split(i, "-")
		intRanges = append(intRanges, intRange{
			low:  utils.IntConv(nums[0]),
			high: utils.IntConv(nums[1]),
		})
	}

	partOneResult := partOne(intRanges)
	partTwoResult := partTwo(intRanges)
	return Result{PartOne: partOneResult, PartTwo: partTwoResult}, nil
}

func partOne(ranges []intRange) int {
	numInvalid := 0
	for _, r := range ranges {
		for i := r.low; i <= r.high; i++ {
			digits := numDigits(i)
			if digits%2 != 0 {
				continue
			}
			l, r := utils.DivMod(i, utils.IntPow(10, digits/2))
			if l == r {
				numInvalid += i
			}
		}
	}
	return numInvalid
}

func partTwo(ranges []intRange) int {
	numInvalid := 0
	for _, r := range ranges {
		for i := r.low; i <= r.high; i++ {

			digits := numDigits(i)
			for s := range digits {
				if s == 0 {
					continue
				}
				if digits%s != 0 {
					continue
				}
				exp := s

				t := i
				parts := make([]int, 0)
				for t >= 1 {
					l := utils.Mod(t, utils.IntPow(10, exp))
					parts = append(parts, l)
					t = utils.DivFloor(t, utils.IntPow(10, exp))
				}
				if len(slices.Compact(parts)) == 1 {
					numInvalid += i
					break
				}
			}
		}
	}
	return numInvalid
}

func numDigits(n int) int {
	digits := 0
	for n >= 1 {
		n /= 10
		digits++
	}
	return digits
}
