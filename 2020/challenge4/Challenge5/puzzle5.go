package challenge5

import (
	"AdventOfCode/file_utils"
	"fmt"
)

var front string = "F"
var left string = "L"

func Run() {
	rows := file_utils.OpenFileIntoSlice("./challenge5/input.txt")
	max := -1
	ticketMap := make(map[int]bool)
	ticketIDs := make([]int, 0)
	for _, row := range rows {
		ticketID, potentialSeat := getTicketIDFromInput(row)
		if ticketID > max {
			max = ticketID
		}
		if potentialSeat {
			ticketMap[ticketID] = true
			ticketIDs = append(ticketIDs, ticketID)
		}
	}
	fmt.Printf("MAX: %d\n", max)
	mySeat := findMySeat(ticketMap, ticketIDs)
	fmt.Printf("My Seat: %d\n", mySeat)
}

func getTicketIDFromInput(row string) (int, bool) {
	potentialSeat := true
	rowString := row[0:7]
	seatString := row[7:]
	rowNum := getNumFromInput(rowString, front, 0, 127)
	seatNum := getNumFromInput(seatString, left, 0, 7)
	ticketID := ((rowNum * 8) + seatNum)
	if rowNum == 0 || rowNum == 127 {
		potentialSeat = false
	}
	return ticketID, potentialSeat
}

func getNumFromInput(input, frontIndicator string, low, high int) int {
	if len(input) == 1 {
		if input == frontIndicator {
			return low
		}
		return high
	}
	mdpt := ((high + low) / 2)
	if input[0:1] == frontIndicator {
		return getNumFromInput(input[1:], frontIndicator, low, mdpt)
	}

	return getNumFromInput(input[1:], frontIndicator, mdpt+1, high)
}

func findMySeat(ticketMap map[int]bool, ticketIDs []int) int {
	var myTicket int
	for _, ticket := range ticketIDs {
		lowRangeTicket, highRangeTicket := ticket-2, ticket+2
		lowTicket, highTicket := ticket-1, ticket+1
		for _, tickets := range [][]int{{lowRangeTicket, lowTicket}, {highRangeTicket, highTicket}} {
			rangeTicket, potentialTicket := tickets[0], tickets[1]
			if _, foundRange := ticketMap[rangeTicket]; foundRange {
				if _, ticketTaken := ticketMap[potentialTicket]; !ticketTaken {
					myTicket = potentialTicket
					break
				}
			}
		}
	}
	return myTicket
}
