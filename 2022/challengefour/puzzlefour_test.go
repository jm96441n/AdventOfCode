package challengefour_test

import (
	"AdventOfCode/2022/challengefour"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengefour.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartOne != 2 {
		t.Errorf("Expected \"2\", got \"%d\"", res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengefour.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != 4 {
		t.Errorf("Expected \"4\", got \"%d\"", res.PartOne)
	}
}
