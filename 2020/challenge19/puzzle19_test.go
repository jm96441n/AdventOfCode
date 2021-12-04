package challenge19_test

import (
	"AdventOfCode/2020/challenge19"
	"testing"
)

func TestFindMatch(t *testing.T) {
	actual := challenge19.FindMatches()
	if actual != 2 {
		t.Errorf("Expected 2, got %d\n", actual)
	}
}
