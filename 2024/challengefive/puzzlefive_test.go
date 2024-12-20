package challengefive_test

import (
	"AdventOfCode/2024/challengefive"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengefive.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 143

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengefive.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 0

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
