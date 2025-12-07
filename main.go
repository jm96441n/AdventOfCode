package main

import (
	"AdventOfCode/2025/challengeseven"
  "flag"
  "log"
)

func main() {
	var isTest bool
	flag.BoolVar(&isTest, "test", false, "run test input")
	flag.Parse()
	inputFile := "./2025/challengeseven/input.txt"
	if isTest {
		inputFile = "./2025/challengeseven/test_input.txt"
	}
  res, err := challengeseven.Run(inputFile)
  if err != nil {
    log.Fatal(err)
  }
  res.Display()
}
