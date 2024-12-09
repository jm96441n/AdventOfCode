package challenge16

import (
	"AdventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type valRange struct {
	high int
	low  int
}

type rule struct {
	name   string
	ranges []valRange
	col    int
}

func (r *rule) validValue(val int) bool {
	valid := false
	for _, vals := range r.ranges {
		if val >= vals.low && val <= vals.high {
			valid = true
		}
	}
	return valid
}

func Run() {
	rows := utils.OpenFileIntoStringSlice("challenge16/input.txt")
	myTicket, otherTickets, rules := parseRows(rows)
	sum, validTickets := sumInvalidTickets(otherTickets, rules)
	fmt.Println(sum)
	fmt.Println(partTwo(validTickets, myTicket, rules))

}

func partTwo(tickets [][]int, myTicket []int, rules []*rule) int {
	//assignedCols := 0

	potentialColMap := make(map[int][]*rule)
	for col := 0; col < len(tickets[0]); col++ {
		potentialColMap[col] = make([]*rule, 0)
	}
	for {
		for col := range potentialColMap {
			for _, rule := range rules {
				if rule.col >= 0 {
					continue
				}
				matchedCol := true
				for row := 0; row < len(tickets); row++ {
					if !rule.validValue(tickets[row][col]) {
						matchedCol = false
						break
					}
				}
				if matchedCol {
					potentialColMap[col] = append(potentialColMap[col], rule)
				}
			}
		}
		for col, colRules := range potentialColMap {
			if len(colRules) == 1 {
				colRules[0].col = col
				delete(potentialColMap, col)
			} else {
				potentialColMap[col] = colRules[:0]
			}
		}
		if len(potentialColMap) == 0 {
			break
		}
	}

	prod := 1
	for _, rule := range rules {
		if strings.Contains(rule.name, "departure") {
			prod *= myTicket[rule.col]
		}
	}
	return prod
}

func sumInvalidTickets(tickets [][]int, rules []*rule) (int, [][]int) {
	invalidMap := make(map[int]struct{})
	validTickets := make([][]int, 0)
	var sum int
	for _, ticket := range tickets {
		valid := true
		for _, num := range ticket {
			invalidCount := 0
			if _, ok := invalidMap[num]; ok {
				sum += num
				valid = false
				break
			} else {
				for _, rule := range rules {
					if !rule.validValue(num) {
						invalidCount++
					}
				}
				if invalidCount == len(rules) {
					sum += num
					valid = false
					break
				}
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	return sum, validTickets
}

func parseRows(rows []string) ([]int, [][]int, []*rule) {
	rules := make([]*rule, 0)
	otherTickets := make([][]int, 0)
	inMyTicket := false
	inNearbyTickets := false
	var myTicket []int
	for _, row := range rows {
		if inMyTicket {
			myTicket = parseTicketNumbers(row)
			inMyTicket = false
		} else if inNearbyTickets {
			otherTickets = append(otherTickets, parseTicketNumbers(row))
		} else if strings.Contains(row, "your ticket") {
			inMyTicket = true
		} else if strings.Contains(row, "nearby tickets") {
			inNearbyTickets = true
		} else if len(row) != 0 {
			rules = append(rules, parseRule(row))
		}
	}
	return myTicket, otherTickets, rules
}

func parseTicketNumbers(row string) []int {
	intTickets := make([]int, 0)
	ticketNums := strings.Split(row, ",")
	for _, ticket := range ticketNums {
		tNum, err := strconv.Atoi(ticket)
		if err != nil {
			panic(err)
		}
		intTickets = append(intTickets, tNum)
	}
	return intTickets
}

var digitRegex = regexp.MustCompile("\\d+")
var identifierRegex = regexp.MustCompile("^(\\w+)(\\s*)(\\w*)")

func parseRule(row string) *rule {
	matches := digitRegex.FindAllString(row, -1)
	valRanges := make([]valRange, 2)
	for i, j := 0, 0; i < len(valRanges) && j < len(matches); i, j = i+1, j+2 {
		low, err := strconv.Atoi(matches[j])
		if err != nil {
			panic(err)
		}

		high, err := strconv.Atoi(matches[j+1])
		if err != nil {
			panic(err)
		}
		valRanges[i] = valRange{
			high: high,
			low:  low,
		}

	}
	name := identifierRegex.FindString(row)
	r := rule{
		name:   name,
		ranges: valRanges,
		col:    -1,
	}
	return &r
}
