package challengeseven_test

import (
	"AdventOfCode/2022/challengeseven"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	res, err := challengeseven.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	if res.PartOne != 95437 {
		t.Errorf("Expected: 95437, Got %d", res.PartOne)
	}
}

func TestRunPartTwo(t *testing.T) {
	res, err := challengeseven.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

	if res.PartTwo != 24933642 {
		t.Errorf("Expected: 24933642, Got %d", res.PartTwo)
	}
}
