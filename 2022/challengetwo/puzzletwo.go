package challengetwo

import (
	"AdventOfCode/2022/file"
	"fmt"
)

type Pair struct {
	Elf string
	Me  string
}

func partOne(plays []Pair, scores map[string]int) int {
	wins := map[string]map[string]int{
		"X": {
			"A": 3,
			"B": 0,
			"C": 6,
		},
		"Y": {
			"A": 6,
			"B": 3,
			"C": 0,
		},
		"Z": {
			"A": 0,
			"B": 6,
			"C": 3,
		},
	}
	score := 0
	for _, p := range plays {
		score += wins[p.Me][p.Elf] + scores[p.Me]
	}

	return score
}

func partTwo(plays []Pair, scores map[string]int) int {
	outcomeVal := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	wins := map[string]map[string]int{
		"X": {
			"A": 3,
			"B": 1,
			"C": 2,
		},
		"Y": {
			"A": 1,
			"B": 2,
			"C": 3,
		},
		"Z": {
			"A": 2,
			"B": 3,
			"C": 1,
		},
	}

	score := 0
	for _, p := range plays {
		fmt.Println(outcomeVal[p.Me])
		fmt.Println(wins[p.Me][p.Elf])
		fmt.Println()

		score += outcomeVal[p.Me] + wins[p.Me][p.Elf]
	}
	return score
}

func Run() {
	scores := map[string]int{
		"X": 1,
		"A": 1,
		"Y": 2,
		"B": 2,
		"Z": 3,
		"C": 3,
	}

	vals := file.OpenFileIntoSlice("./challengetwo/input.txt", file.StringConv)
	plays := make([]Pair, 0)
	for _, v := range vals {
		plays = append(plays, Pair{Elf: string(v[0]), Me: string(v[len(v)-1])})
	}
	pOne := partOne(plays, scores)
	pTwo := partTwo(plays, scores)
	fmt.Printf("Part One: %d\nPart Two: %d\n", pOne, pTwo)
}
