package challengefive

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

type Range struct {
	high int
	low  int
}

func (r Range) includes(val int) bool {
	return val >= r.low && val <= r.high
}

func (r Range) valGreaterThan(val int) bool {
	return val > r.high
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.low, r.high)
}

func Run(filename string) (Result, error) {
	vals := utils.OpenFileIntoSlice(filename, utils.StringConv)
	ranges := make([]Range, 0, len(vals))
	ids := make([]int, 0, len(vals))
	inRanges := true
	for _, v := range vals {
		if len(v) == 0 {
			inRanges = false
			continue
		}

		if inRanges {
			nums := strings.Split(v, "-")
			ranges = append(ranges, Range{
				high: utils.IntConv(nums[1]),
				low:  utils.IntConv(nums[0]),
			})
		} else {
			ids = append(ids, utils.IntConv(v))
		}
	}
	return Result{
		PartOne: partOne(ranges, ids),
		PartTwo: partTwo(ranges),
	}, nil
}

func partOne(ranges []Range, ids []int) int {
	sum := 0

	mergedRanges := sortAndMergeRanges(ranges)
	numRanges := len(mergedRanges)
	for _, i := range ids {
		if bSearch(mergedRanges, i, 0, numRanges-1) {
			sum += 1
		}
	}
	return sum
}

func partTwo(ranges []Range) int {
	n := 0
	mergedRanges := sortAndMergeRanges(ranges)
	for _, r := range mergedRanges {
		n += (r.high - r.low + 1)
	}
	return n
}

func sortAndMergeRanges(ranges []Range) []Range {
	slices.SortFunc(ranges, func(a, b Range) int {
		if a.high <= b.low {
			return -1
		}

		if b.high <= a.low {
			return 1
		}

		if a.low < b.low && a.high > b.low {
			return -1
		}

		if b.low < a.low && b.high > a.low {
			return 1
		}

		return 0
	})

	mergedRanges := make([]Range, 0, len(ranges))
	j := 1
	i := 0
	for i < len(ranges)-1 {
		newRange := Range{low: ranges[i].low, high: ranges[i].high}
		for j < len(ranges) && (ranges[j].low <= newRange.high) {
			newRange.high = max(newRange.high, ranges[j].high)
			j++
		}
		i = j
		mergedRanges = append(mergedRanges, newRange)
	}
	return mergedRanges
}

func bSearch(ranges []Range, id int, left, right int) bool {
	if left > right {
		return false
	}

	mdpt := left + (utils.DivFloor((right - left), 2))
	if ranges[mdpt].includes(id) {
		return true
	}

	if ranges[mdpt].valGreaterThan(id) {
		return bSearch(ranges, id, mdpt+1, right)
	}

	return bSearch(ranges, id, left, mdpt-1)
}
