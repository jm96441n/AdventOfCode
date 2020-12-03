package file_utils

import (
	"bufio"
	"log"
	"os"
)

func OpenFileIntoSlice(fileName string) []string {
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
	return rows
}
