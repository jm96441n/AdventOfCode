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
	degreeMap := buildDegreeMap()
	orientationMap := buildOrientationMap()
	for _, a := range actions {
		if a.direction == RIGHT || a.direction == LEFT {
			s.orientation = getNewOrientation(s, a, orientationMap, degreeMap)
		} else if a.direction == FORWARD {
			moveDirection(s, s.orientation, a.distance)
		} else {
			moveDirection(s, a.direction, a.distance)
		}
	}
}

func moveDirection(s *ship, direction string, distance int) {
	switch direction {
	case NORTH:
		s.vert += distance
	case SOUTH:
		s.vert -= distance
	case EAST:
		s.horz += distance
	case WEST:
		s.horz -= distance
	}
}

func getNewOrientation(s *ship, a action, oMap map[string]int, degreeMap map[int]string) string {
	curDir := oMap[s.orientation]
	if a.direction == RIGHT {
		curDir += a.distance
		for curDir >= 360 {
			curDir -= 360
		}
	} else if a.direction == LEFT {
		curDir -= a.distance
		for curDir < 0 {
			curDir += 360
		}
	}
	return degreeMap[curDir]
}

func buildDegreeMap() map[int]string {
	turnMap := make(map[int]string)
	turnMap[0] = NORTH
	turnMap[90] = EAST
	turnMap[180] = SOUTH
	turnMap[270] = WEST
	return turnMap
}

func buildOrientationMap() map[string]int {
	oMap := make(map[string]int)
	oMap[NORTH] = 0
	oMap[EAST] = 90
	oMap[SOUTH] = 180
	oMap[WEST] = 270
	return oMap
}
