package challenge14

import (
	"AdventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var digitRegex = regexp.MustCompile("\\d+")

func Run() {
	rows := utils.OpenFileIntoStringSlice("challenge14/input.txt")
	fmt.Println(partOne(rows))
	fmt.Println(partTwo(rows))
}

func partOne(rows []string) int {
	mem := make(map[int]int)
	var (
		mask       string
		memVal     int
		memAddress int
		sum        int
	)
	for _, row := range rows {
		if strings.Contains(row, "mask") {
			fmt.Sscanf(row, "mask = %s", &mask)
		} else {
			memAddress, memVal = applyMaskToRow(row, mask)
			mem[memAddress] = memVal
		}

	}
	for _, val := range mem {
		sum += val
	}
	return sum
}

func applyMaskToRow(row, mask string) (int, int) {
	var (
		memVal     int
		memAddress int
		err        error
	)
	matches := digitRegex.FindAllString(row, -1)

	for idx, val := range []*int{&memAddress, &memVal} {
		*val, err = strconv.Atoi(matches[idx])
		if err != nil {
			panic(err)
		}
	}
	binMemVal := convertToBinary(memVal, mask)
	maskedMemVal := applyMask(binMemVal, mask)
	newMemVal := convertToDecimal(maskedMemVal)
	return memAddress, newMemVal
}

func applyMask(binMemVal, mask string) string {
	maskedBinMemVal := []rune(binMemVal)
	for idx, char := range mask {
		if char != 'X' {
			maskedBinMemVal[idx] = char
		}
	}
	return string(maskedBinMemVal)
}

func partTwo(rows []string) int {
	mem := make(map[int]int)
	var (
		mask         string
		memVal       int
		memAddresses []int
		sum          int
	)
	for _, row := range rows {
		if strings.Contains(row, "mask") {
			fmt.Sscanf(row, "mask = %s", &mask)
		} else {
			memAddresses, memVal = applyVariableMaskToRow(row, mask)
			for _, address := range memAddresses {
				mem[address] = memVal
			}
		}

	}
	for _, val := range mem {
		sum += val
	}
	return sum
}

func applyVariableMaskToRow(row, mask string) ([]int, int) {

	var (
		memVal     int
		memAddress int
		err        error
	)
	matches := digitRegex.FindAllString(row, -1)
	newMemAddresses := make([]int, 0)
	for idx, val := range []*int{&memAddress, &memVal} {
		*val, err = strconv.Atoi(matches[idx])
		if err != nil {
			panic(err)
		}
	}
	binMemAddress := convertToBinary(memAddress, mask)
	maskedBinMemAddress := applyVariableMask(binMemAddress, mask)
	newBinMemAddresses := calcVariableMemAddresses(maskedBinMemAddress)
	for _, newBinMemAddress := range newBinMemAddresses {
		decAdd := convertToDecimal(newBinMemAddress)
		newMemAddresses = append(newMemAddresses, decAdd)
	}
	return newMemAddresses, memVal
}

func applyVariableMask(binMemVal, mask string) []rune {
	maskedBinMemVal := []rune(binMemVal)
	for idx, char := range mask {
		if char != '0' {
			maskedBinMemVal[idx] = char
		}
	}
	return maskedBinMemVal
}

func calcVariableMemAddresses(memVal []rune) []string {
	stringifiedMem := string(memVal)
	if !strings.Contains(stringifiedMem, "X") {
		return []string{stringifiedMem}
	}
	mem1 := make([]rune, len(memVal))
	mem2 := make([]rune, len(memVal))
	copy(mem1, memVal)
	copy(mem2, memVal)

	mems := make([]string, 0)
	for idx, char := range memVal {
		if char == 'X' {
			mem1[idx] = '0'
			mem2[idx] = '1'
			break
		}
	}
	mems = append(mems, calcVariableMemAddresses(mem1)...)
	mems = append(mems, calcVariableMemAddresses(mem2)...)
	return mems
}

func convertToBinary(memVal int, mask string) string {
	var bit int
	var bitString string
	for memVal > 0 {
		memVal, bit = divMod(memVal, 2)
		bitString = fmt.Sprintf("%d%s", bit, bitString)
	}
	for len(bitString) != len(mask) {
		bitString = fmt.Sprintf("%d%s", 0, bitString)
	}
	return bitString
}

func convertToDecimal(binString string) int {
	var decVal int
	for idx, char := range binString {
		pow := (len(binString) - 1) - idx
		if char == '1' {
			decVal += int(math.Pow(2, float64(pow)))
		}
	}
	return decVal
}

func divMod(num, divisor int) (int, int) {
	q := num / divisor
	mod := num % divisor
	return q, mod
}
