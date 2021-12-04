package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	frequency := 0
	frequencies := make(map[int]int)
	vals := make([]int, 0)
	file, err := os.Open("1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		vals = append(vals, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	i := 0
	valCount := len(vals)
	var idx int
	for {
		idx = i % valCount
		frequency += vals[idx]
		if _, ok := frequencies[frequency]; ok {
			frequencies[frequency]++
			break
		} else {
			frequencies[frequency] = 1
		}
		i++
	}

	fmt.Printf("%d\n", frequency)

}
