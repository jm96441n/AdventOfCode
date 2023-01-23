package challengeeight_test

import (
	"AdventOfCode/2022/challengeeight"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengeeight.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartOne != 21 {
		t.Errorf("Expected: 21, Got: %d", res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengeeight.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != 0 {
		t.Errorf("Expected: 0, Got: %d", res.PartTwo)
	}
}
