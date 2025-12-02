package main

import (
	"flag"
	"log"

	"AdventOfCode/2025/challengetwo"
)

func main() {
	var isTest bool
	flag.BoolVar(&isTest, "test", false, "run test input")
	flag.Parse()
	inputFile := "./2025/challengetwo/input.txt"
	if isTest {
		inputFile = "./2025/challengetwo/test_input.txt"
	}
	res, err := challengetwo.Run(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	res.Display()
}
