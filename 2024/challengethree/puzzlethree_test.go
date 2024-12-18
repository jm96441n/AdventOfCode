package challengethree_test

import (
	"AdventOfCode/2024/challengethree"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengethree.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expectedValue := 161

	if res.PartOne != expectedValue {
		t.Errorf("Expected: %d, Got: %d", expectedValue, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengethree.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expectedValue := 48

	if res.PartTwo != expectedValue {
		t.Errorf("Expected: %d, Got: %d", expectedValue, res.PartOne)
	}
}
