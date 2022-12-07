package challengesix_test

import (
	"AdventOfCode/2022/challengesix"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengesix.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartOne != 7 {
		t.Errorf("Expected 4, got %d", res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengesix.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != 19 {
		t.Errorf("Expected 19, got %d", res.PartTwo)
	}
}
