package challenge13

import (
	"AdventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	rows := utils.OpenFileIntoStringSlice("challenge13/input.txt")
	earliestTime, err := strconv.Atoi(rows[0])
	if err != nil {
		panic(err)
	}
	busses := strings.Split(rows[1], ",")
	fmt.Println(partOneProduct(earliestTime, busses))
	fmt.Println(partTwoTime(busses))
}

func partOneProduct(earliestTime int, busses []string) int {
	busNumbers := make([]int, 0)
	for _, bus := range busses {
		if bus != "x" {
			num, err := strconv.Atoi(bus)
			if err != nil {
				panic(err)
			}
			busNumbers = append(busNumbers, num)
		}
	}
	earliestBus, earliestBusDeparture := calcEarliestBus(earliestTime, busNumbers)
	return (earliestBusDeparture - earliestTime) * earliestBus
}

func calcEarliestBus(earliestTime int, busNumbers []int) (int, int) {
	var (
		divisor              int
		departure            int
		earliestBusDeparture int
		earliestBus          int
	)
	for _, bus := range busNumbers {
		divisor = earliestTime / bus
		departure = bus * (divisor + 1)
		if departure < earliestBusDeparture || earliestBusDeparture == 0 {
			earliestBusDeparture = departure
			earliestBus = bus
		}
	}
	return earliestBus, earliestBusDeparture
}

func partTwoTime(busses []string) int {
	busNumbers := make([]int, 0)
	offsets := make([]int, 0)
	for idx, bus := range busses {
		if bus != "x" {
			num, err := strconv.Atoi(bus)
			if err != nil {
				panic(err)
			}
			offsets = append(offsets, idx)
			busNumbers = append(busNumbers, num)
		}
	}
	timest := 1
	for {
		timeSkip := 1
		validTime := true
		for idx, offset := range offsets {
			if (timest+offset)%busNumbers[idx] != 0 {
				validTime = false
				break
			}
			timeSkip *= busNumbers[idx]
		}
		if validTime {
			break
		}
		timest += timeSkip
	}
	return timest
}
