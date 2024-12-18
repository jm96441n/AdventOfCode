package challengefour_test

import (
	"AdventOfCode/2024/challengefour"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengefour.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 18

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengefour.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 0

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
