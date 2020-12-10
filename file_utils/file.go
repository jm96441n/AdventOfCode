package file_utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func OpenFileIntoSlice(fileName string) []string {
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
