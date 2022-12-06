package challengefive_test

import (
	"AdventOfCode/2022/challengefive"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengefive.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartOne != "CMZ" {
		t.Errorf("Expected \"CMZ\", got %q", res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengefive.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != "MCD" {
		t.Errorf("Expected \"MCD\", got %q", res.PartTwo)
	}
}
