package challengethree

import (
	"AdventOfCode/2022/file"
	"fmt"
)

func Run() {
	sacks := file.OpenFileIntoStringSlice("./challengethree/input.txt", file.StringConv)
	fmt.Printf("Part One: %d\nPart Two: %d\n", partOne(sacks), partTwo(sacks))
}

func partOne(sacks []string) int {
	sum := 0
	for _, sack := range sacks {
		var matching byte

		items := itemsToSet(sack[0 : len(sack)/2])

		for j := len(sack) / 2; j < len(sack); j++ {
			_, ok := items[sack[j]]
			if ok {
				matching = sack[j]
				break

			}
		}
		sum += priority(matching)
	}
	return sum
}

func partTwo(sacks []string) int {
	sum := 0

	for i := 0; i < len(sacks)-2; i += 3 {
		var matching byte
		r3 := sacks[i+2]

		itemsOne := itemsToSet(sacks[i])
		itemsTwo := itemsToSet(sacks[i+1])

		for j := 0; j < len(r3); j += 1 {
			_, inOne := itemsOne[r3[j]]
			_, inTwo := itemsTwo[r3[j]]
			if inOne && inTwo {
				matching = r3[j]
				break
			}
		}
		sum += priority(matching)
	}
	return sum
}

func itemsToSet(sack string) map[byte]struct{} {
	items := make(map[byte]struct{}, 0)
	for i := 0; i < len(sack); i++ {
		items[sack[i]] = struct{}{}
	}
	return items
}

func priority(matching byte) int {
	if int(matching) >= 97 && int(matching) <= 122 {
		return int(matching) - int('a') + 1
	}
	return int(matching) - int('A') + 27
}
