package file

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type sliceEle interface {
	string | int
}

func OpenFileIntoSlice[T sliceEle](fileName string, conv func(string) T) []T {
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
