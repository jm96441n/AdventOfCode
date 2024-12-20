package challengefive

import (
	"AdventOfCode/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
	input := utils.OpenFileIntoSlice[string](filename, utils.StringConv)
	pages, updates := parseInput(input)

	p1 := partOne(pages, updates)
	return Result{
		PartOne: p1,
	}, nil
}

func parseInput(input []string) ([][2]int, [][]int) {
	pages := make([][2]int, 0)
	updates := make([][]int, 0)
	first := make([]int, 0)
	second := make([]int, 0)
	for _, line := range input {
		if strings.Contains(line, "|") {
			splits := strings.Split(line, "|")
			d1, err := strconv.Atoi(splits[0])
			if err != nil {
				panic(err)
			}
			first = append(first, d1)

			d2, err := strconv.Atoi(splits[1])
			if err != nil {
				panic(err)
			}
			second = append(second, d2)

			pages = append(pages, [2]int{d1, d2})
		} else if strings.Contains(line, ",") {
			splits := strings.Split(line, ",")
			update := make([]int, 0, len(splits))

			for _, s := range splits {
				d, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				update = append(update, d)
			}
			updates = append(updates, update)
		}
	}
	slices.Sort(first)
	slices.Sort(second)
	fmt.Println(first)
	fmt.Println(second)

	return pages, updates
}

func partOne(pages [][2]int, updates [][]int) int {
	sum := 0

	graph := make(map[int]map[int]struct{}, 0)

	for _, pair := range pages {
		if _, ok := graph[pair[1]]; !ok {
			graph[pair[1]] = make(map[int]struct{}, 0)
		}
		graph[pair[1]][pair[0]] = struct{}{}
	}

	return sum
}

func topoSort(pages [][2]int) []int {
	sortedEles := make([]int, 0, len(pages)*2)

	elesWithNotIncomingEdge := make([]int, 0)

	graph := make(map[int]map[int]struct{}, 0)

	for _, pair := range pages {
		if _, ok := graph[pair[1]]; !ok {
			graph[pair[1]] = make(map[int]struct{}, 0)
		}
		graph[pair[1]][pair[0]] = struct{}{}
	}

	for k, v := range graph {
		fmt.Println(k, v)
		if len(v) == 0 {
			elesWithNotIncomingEdge = append(elesWithNotIncomingEdge, k)
		}
	}

	for len(elesWithNotIncomingEdge) > 0 {
		eleToRemove := elesWithNotIncomingEdge[0]
		elesWithNotIncomingEdge = elesWithNotIncomingEdge[1:]
		for k, v := range graph {
			delete(v, eleToRemove)
			if len(v) == 0 {
				elesWithNotIncomingEdge = append(elesWithNotIncomingEdge, k)
				delete(graph, k)
				continue
			}
			graph[k] = v
		}
	}

	if len(graph) != 0 {
		fmt.Println(graph)
		fmt.Println(sortedEles)
		panic("expected graph to be empty, still had elements")
	}

	return sortedEles
}
