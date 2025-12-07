package challengesix_test

import (
	"AdventOfCode/2025/challengesix"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengesix.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 4277556

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengesix.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 3263827

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
