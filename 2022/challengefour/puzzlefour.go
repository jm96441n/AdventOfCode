package challengefour

import (
	"AdventOfCode/2022/file"
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

type WorkRange struct {
	Start int
	End   int
}

func rangeFromString(in string) (WorkRange, error) {
	ranges := strings.Split(in, "-")
	start, err := strconv.Atoi(ranges[0])
	if err != nil {
		return WorkRange{}, err
	}
	end, err := strconv.Atoi(ranges[1])
	if err != nil {
		return WorkRange{}, err
	}

	return WorkRange{
		Start: start,
		End:   end,
	}, nil
}

func (w WorkRange) contains(w2 WorkRange) bool {
	return w.Start <= w2.Start && w.End >= w2.End
}

func (w WorkRange) overlaps(w2 WorkRange) bool {
	return (w.Start <= w2.Start && w2.Start <= w.End) || (w.End >= w2.End && w.Start <= w2.End)
}

type ElfPair struct {
	One WorkRange
	Two WorkRange
}

func Run(filename string) (Result, error) {
	input := file.OpenFileIntoSlice(filename, file.StringConv)
	pairs := make([]ElfPair, len(input))
	for i, work := range input {
		elfs := strings.Split(work, ",")
		elfOne, err := rangeFromString(elfs[0])
		if err != nil {
			return Result{}, err
		}
		elfTwo, err := rangeFromString(elfs[1])
		if err != nil {
			return Result{}, err
		}

		pairs[i] = ElfPair{
			One: elfOne,
			Two: elfTwo,
		}
	}
	totalOverlap := 0
	overlap := 0

	for _, pair := range pairs {
		if pair.One.contains(pair.Two) || pair.Two.contains(pair.One) {
			totalOverlap += 1
		}
		if pair.One.overlaps(pair.Two) || pair.Two.overlaps(pair.One) {
			overlap += 1
		}
	}
	return Result{
		PartOne: totalOverlap,
		PartTwo: overlap,
	}, nil
}
