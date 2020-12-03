package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	ids := make([]string, 0)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	genCheckSum(ids)
	findSimilarIds(ids)
}

func genCheckSum(ids []string) {
	var (
		twice  = 0
		thrice = 0
	)

	for _, id := range ids {
		appearsTwice, appearsThrice := appearances(id)
		if appearsTwice {
			twice++
		}
		if appearsThrice {
			thrice++
		}
	}

	fmt.Println(twice * thrice)
}

func appearances(id string) (bool, bool) {
	var (
		appearsTwice  = false
		appearsThrice = false
	)

	freqs := make(map[rune]int)

	for _, char := range id {
		if _, ok := freqs[char]; ok {
			freqs[char]++
		} else {
			freqs[char] = 1
		}
	}

	for _, count := range freqs {
		if count == 2 {
			appearsTwice = true
		}
		if count == 3 {
			appearsThrice = true
		}
		if appearsTwice && appearsThrice {
			break
		}
	}
	return appearsTwice, appearsThrice
}

func findSimilarIds()
