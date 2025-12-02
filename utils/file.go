package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func OpenFileIntoSlice[T any](fileName string, conv func(string) T) []T {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	rows := make([]T, 0)
	for scanner.Scan() {
		rows = append(rows, conv(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rows
}

func OpenFileIntoString(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func StringConv(s string) string {
	return s
}

func IntConv(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func IntRowConv(s string) []int {
	splits := strings.Fields(s)
	intRow := make([]int, 0, len(splits))
	for _, split := range splits {
		i, err := strconv.Atoi(split)
		if err != nil {
			log.Fatal(err)
		}
		intRow = append(intRow, i)
	}
	return intRow
}

func OpenFileIntoStringSlice(fileName string) []string {
	file, err := os.Open(fileName)
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
	return rows
}

func OpenFileIntoIntSlice(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	rows := make([]int, 0)
	for scanner.Scan() {
		intRow, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		rows = append(rows, intRow)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return rows
}
