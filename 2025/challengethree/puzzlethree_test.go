package challengethree_test

import (
	"testing"

	"AdventOfCode/2025/challengethree"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengethree.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 357

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengethree.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 3121910778619

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
