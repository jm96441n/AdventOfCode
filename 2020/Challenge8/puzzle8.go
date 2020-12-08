package main

import (
	"AdventOfCode/file_utils"
	"fmt"
)

var (
	nop   = "nop"
	acc   = "acc"
	jmp   = "jmp"
	plus  = "+"
	minus = "-"
)

type instruction struct {
	cmd   string
	value int
}

func main() {
	rows := file_utils.OpenFileIntoSlice("input.txt")
	fmt.Println(partOne(rows))
	fmt.Println(partTwo(rows))
}

func partOne(rows []string) int {
	instructions, insMap := buildInstructionsFromRows(rows)
	_, accCounter := run(instructions, insMap)
	return accCounter
}

func partTwo(rows []string) int {
	instructions, insMap := buildInstructionsFromRows(rows)
	var (
		terminated bool
		accCounter int
		prev_cmd   string
	)
	for _, ins := range instructions {
		prev_cmd = ins.cmd
		switch ins.cmd {
		case jmp:
			ins.cmd = nop
		case nop:
			ins.cmd = jmp
		default:
			continue
		}
		terminated, accCounter = run(instructions, insMap)
		if terminated {
			break
		}
		ins.cmd = prev_cmd
		for key, _ := range insMap {
			insMap[key] = 0
		}
	}
	return accCounter
}

func run(instructions []*instruction, insMap map[*instruction]int) (bool, int) {
	accCounter := 0
	i := 0
	terminated := false
	for {
		if i >= len(instructions) {
			terminated = true
			break
		}
		ins := instructions[i]
		if val, ok := insMap[ins]; ok && val > 0 {
			break
		}

		switch ins.cmd {
		case acc:
			{
				accCounter += ins.value
				insMap[ins]++
				i++
			}
		case nop:
			i++
		case jmp:
			i += ins.value
		}
	}
	return terminated, accCounter
}

func buildInstructionsFromRows(rows []string) ([]*instruction, map[*instruction]int) {
	instructions := make([]*instruction, len(rows))
	insMap := make(map[*instruction]int)
	var (
		cmd   string
		value int
	)
	for idx, row := range rows {
		fmt.Sscanf(row, "%s %d", &cmd, &value)
		ins := &instruction{
			cmd:   cmd,
			value: value,
		}
		instructions[idx] = ins
		insMap[ins] = 0
	}
	return instructions, insMap
}
