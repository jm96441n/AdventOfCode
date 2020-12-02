package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type PasswordValidation struct {
	high     int
	low      int
	letter   string
	password string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	rows := make([]string, 0)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	allSplits := buildSplits(rows)
	partOneCount := partOne(allSplits)
	partTwoCount := partTwo(allSplits)
	fmt.Println(partOneCount)
	fmt.Println(partTwoCount)
}

func partOne(allSplits []PasswordValidation) int {
	validCount := 0
	for _, split := range allSplits {
		matchCount := strings.Count(split.password, split.letter)
		if matchCount >= split.low && matchCount <= split.high {
			validCount++
		}
	}
	return validCount
}

func partTwo(allSplits []PasswordValidation) int {
	validCount := 0
	for _, split := range allSplits {
		first := letterMatches(split.password, split.letter, split.low)
		second := letterMatches(split.password, split.letter, split.high)
		if first != second {
			validCount++
		}
	}
	return validCount
}

func letterMatches(password string, letter string, count int) bool {
	idx := count - 1
	if idx >= len(password) || idx < 0 {
		return false
	}
	return string(password[idx]) == letter
}

func buildSplits(rows []string) []PasswordValidation {
	allSplits := make([]PasswordValidation, 0)
	re := regexp.MustCompile(":?[[:space:]]")
	for _, row := range rows {
		splits := re.Split(row, -1)
		freqs, letter, pass := splits[0], strings.Split(splits[1], ":")[0], splits[2]
		freqSplits := strings.Split(freqs, "-")
		low, err := strconv.Atoi(freqSplits[0])
		if err != nil {
			log.Fatal(err)
		}

		high, err := strconv.Atoi(freqSplits[1])
		if err != nil {
			log.Fatal(err)
		}
		split := PasswordValidation{
			high:     high,
			low:      low,
			letter:   letter,
			password: pass,
		}
		allSplits = append(allSplits, split)
	}
	return allSplits
}
