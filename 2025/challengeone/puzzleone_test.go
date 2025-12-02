package challengeone_test

import (

	"AdventOfCode/2025/challengeone"
	"testing"
)

func TestRunPartOne(t *testing.T) {
  res, err := challengeone.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

  expected := 0

	if res.PartOne != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartOne)
	}

}
func TestRunPartTwo(t *testing.T) {
  res, err := challengeone.Run("./test_input.txt")
	if err != nil {
		t.Error(err)
	}

  expected := 0

	if res.PartTwo != expected {
		t.Errorf("Expected: %d, Got: %d", expected, res.PartTwo)
	}

}
