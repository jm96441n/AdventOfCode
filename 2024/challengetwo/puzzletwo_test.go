package challengetwo_test

import (
	"AdventOfCode/2024/challengetwo"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengetwo.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartOne != 2 {
		t.Errorf("Expected: 2, Got: %d", res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengetwo.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != 4 {
		t.Errorf("Expected: 4, Got: %d", res.PartTwo)
	}
}
