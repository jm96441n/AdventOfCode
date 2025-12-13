package challengeeight_test

import (
	"AdventOfCode/2025/challengeeight"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengeeight.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 40

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengeeight.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 0

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
