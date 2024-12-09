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
	if res.PartOne != 0 {
		t.Errorf("Expected: 0, Got: %d", res.PartOne)
	}

}
func TestRunPartTwo(t *testing.T) {
  res, err := challengethree.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != 0 {
		t.Errorf("Expected: 0, Got: %d", res.PartTwo)
	}

}
