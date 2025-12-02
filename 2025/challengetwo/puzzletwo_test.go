package challengetwo_test

import (
	"testing"

	"AdventOfCode/2025/challengetwo"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengetwo.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 1227775554

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengetwo.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 4174379265

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}
}
