package challengeone_test

import (
	"AdventOfCode/2024/challengeone"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengeone.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartOne != 11 {
		t.Errorf("Expected: 11, Got: %d", res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengeone.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != 31 {
		t.Errorf("Expected: 31, Got: %d", res.PartTwo)
	}
}
