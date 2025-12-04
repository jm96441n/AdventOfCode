package challengefour_test

import (
	"testing"

	"AdventOfCode/2025/challengefour"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengefour.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 13

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengefour.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 43

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
