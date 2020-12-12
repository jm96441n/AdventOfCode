package challenge12

import (
	"AdventOfCode/file_utils"
	"fmt"
	"regexp"
	"strconv"
)

type action struct {
	direction string
	distance  int
}

type ship struct {
	orientation string
	horz        int
	vert        int
}

var (
	EAST    = "E"
	WEST    = "W"
	NORTH   = "N"
	SOUTH   = "S"
	RIGHT   = "R"
	LEFT    = "L"
	FORWARD = "F"
)

func Run() {
	rows := file_utils.OpenFileIntoSlice("challenge12/input.txt")
	formattedRows := formatRows(rows)
	md := calculateManhattanDistance(formattedRows)
	fmt.Println(md)
}

func formatRows(rows []string) []action {
	formattedRows := make([]action, len(rows))
	var (
		movement  int
		direction string
		err       error
	)
	re1 := regexp.MustCompile("^(\\w){1}")
	re2 := regexp.MustCompile("(\\d)+$")
	for idx, row := range rows {
		direction = re1.FindStringSubmatch(row)[0]
		movement, err = strconv.Atoi(re2.FindStringSubmatch(row)[0])
		if err != nil {
			panic(err)
		}
		formattedRows[idx] = action{
			direction: direction,
			distance:  movement,
		}
	}
	return formattedRows
}

func calculateManhattanDistance(actions []action) int {
	s := &ship{orientation: "E"}
	runTheCourse(s, actions)
	return abs(s.horz) + abs(s.vert)
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func runTheCourse(s *ship, actions []action) {
	turnMap := buildTurnMap()
	for _, a := range actions {
		switch a.direction {
		case NORTH:
			s.vert += a.distance
		case SOUTH:
			s.vert -= a.distance
		case EAST:
			s.horz += a.distance
		case WEST:
			s.horz -= a.distance
		case RIGHT:
			s.orientation = turnMap[s.orientation][((-1) * (a.distance))]
		case LEFT:
			s.orientation = turnMap[s.orientation][a.distance]
		case FORWARD:
			switch s.orientation {
			case NORTH:
				s.vert += a.distance
			case SOUTH:
				s.vert -= a.distance
			case EAST:
				s.horz += a.distance
			case WEST:
				s.horz -= a.distance
			}
		default:
			panic(fmt.Sprintf("UNKNOWN DIRECTION: %s", a.direction))
		}
	}
}

func buildTurnMap() map[string]map[int]string {
	// assume that a left turn is positive, a right turn is negative
	turnMap := make(map[string]map[int]string)
	turnMap[EAST] = make(map[int]string)
	turnMap[WEST] = make(map[int]string)
	turnMap[SOUTH] = make(map[int]string)
	turnMap[NORTH] = make(map[int]string)

	turnMap[EAST][90] = NORTH
	turnMap[EAST][180] = WEST
	turnMap[EAST][270] = SOUTH

	turnMap[EAST][-90] = SOUTH
	turnMap[EAST][-180] = WEST
	turnMap[EAST][-270] = NORTH

	turnMap[WEST][90] = SOUTH
	turnMap[WEST][180] = EAST
	turnMap[WEST][270] = NORTH

	turnMap[WEST][-180] = EAST
	turnMap[WEST][-90] = NORTH
	turnMap[WEST][-270] = SOUTH

	turnMap[NORTH][90] = WEST
	turnMap[NORTH][180] = SOUTH
	turnMap[NORTH][270] = EAST

	turnMap[NORTH][-90] = EAST
	turnMap[NORTH][-180] = SOUTH
	turnMap[NORTH][-270] = WEST

	turnMap[SOUTH][90] = EAST
	turnMap[SOUTH][180] = NORTH
	turnMap[SOUTH][270] = WEST

	turnMap[SOUTH][-90] = WEST
	turnMap[SOUTH][-180] = NORTH
	turnMap[SOUTH][-270] = EAST
	return turnMap
}
