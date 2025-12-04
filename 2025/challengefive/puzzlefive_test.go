package challengefive_test

import (
	"testing"

	"AdventOfCode/2025/challengefive"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengefive.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 3

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengefive.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 14

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
