package challengeeight

import (
	"AdventOfCode/utils"
	"fmt"
	"math"
	"strings"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

type point struct {
	x int
	y int
	z int
}

func (p point) distance(other point) float64 {
	return math.Sqrt(
		float64(utils.IntPow(p.x-other.x, 2) + utils.IntPow(p.y-other.y, 2) + utils.IntPow(p.z-other.z, 2)),
	)
}

func Pointconv(s string) point {
	parts := strings.Split(s, ",")
	return point{
		x: utils.IntConv(parts[0]),
		y: utils.IntConv(parts[1]),
		z: utils.IntConv(parts[2]),
	}
}

func Run(filename string) (Result, error) {
	points := utils.OpenFileIntoSlice(filename, Pointconv)
	return Result{
		PartOne: partOne(points),
	}, nil
}

func partOne(points []point) int {
	q := utils.PriorityQueue[float64]{}
	distToPoint := make(map[float64][2]point)

	for i := 0; i < len(points); i++ {
		start := points[i]
		for j := i + 1; j < len(points); j++ {
			end := points[j]
			dist := start.distance(end)
			q.Push(dist)
			distToPoint[dist] = [2]point{start, end}
		}
	}

	clusters := make([]map[point]struct{}, 0)
	for i := 0; i <= 1000; i++ {
		d := q.Pop()
		connected := distToPoint[d]
		added := false
		for _, cluster := range clusters {
			if _, ok := cluster[connected[0]]; ok {
				cluster[connected[1]] = struct{}{}
				added = true
				break
			}
			if _, ok := cluster[connected[1]]; ok {
				cluster[connected[0]] = struct{}{}
				added = true
				break
			}
		}
		if !added {
			clusters = append(clusters, map[point]struct{}{
				connected[0]: {},
				connected[1]: {},
			})
		}
	}

	pq := utils.PriorityQueue[int]{}
	for _, c := range clusters {
		fmt.Println(len(c))
		pq.Push(len(c))
		if pq.Size > 3 {
			pq.Pop()
		}
	}
	size := 1
	for range 3 {
		clSize := pq.Pop()
		fmt.Println(clSize)
		size *= clSize
	}
	return size
}
