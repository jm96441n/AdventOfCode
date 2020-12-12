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

type waypoint struct {
	horz int
	vert int
}

type courseRunner func(s *ship, actions []action)

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
	for _, runner := range []courseRunner{runTheCoursePartOne, runTheCoursePartTwo} {
		fmt.Println(calculateManhattanDistance(formattedRows, runner))
	}
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

func calculateManhattanDistance(actions []action, runTheCourse courseRunner) int {
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

func runTheCoursePartOne(s *ship, actions []action) {
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

func runTheCoursePartTwo(s *ship, actions []action) {
	wp := waypoint{
		horz: 10,
		vert: 1,
	}
	for _, a := range actions {
		if a.direction == RIGHT {
			wp = rotateWaypoint(wp, (360 - a.distance))
		} else if a.direction == LEFT {
			wp = rotateWaypoint(wp, a.distance)
		} else if a.direction == FORWARD {
			s.vert += (a.distance * wp.vert)
			s.horz += (a.distance * wp.horz)
		} else {
			wp = moveWaypoint(wp, a)
		}
	}
}

func rotateWaypoint(wp waypoint, rotation int) waypoint {
	var movedWp waypoint
	switch rotation {
	case 90:
		movedWp.horz = (-1) * wp.vert
		movedWp.vert = wp.horz
	case 180:
		movedWp.horz = (-1) * wp.horz
		movedWp.vert = (-1) * wp.vert
	case 270:
		movedWp.horz = wp.vert
		movedWp.vert = (-1) * wp.horz
	}
	return movedWp
}

func moveWaypoint(wp waypoint, a action) waypoint {
	movedWp := waypoint{
		horz: wp.horz,
		vert: wp.vert,
	}
	switch a.direction {
	case NORTH:
		movedWp.vert += a.distance
	case SOUTH:
		movedWp.vert -= a.distance
	case EAST:
		movedWp.horz += a.distance
	case WEST:
		movedWp.horz -= a.distance
	}
	return movedWp
}
