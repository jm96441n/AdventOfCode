package main

import (
	"AdventOfCode/2025/challengefour"
  "flag"
  "log"
)

func main() {
	var isTest bool
	flag.BoolVar(&isTest, "test", false, "run test input")
	flag.Parse()
	inputFile := "./2025/challengefour/input.txt"
	if isTest {
		inputFile = "./2025/challengefour/test_input.txt"
	}
  res, err := challengefour.Run(inputFile)
  if err != nil {
    log.Fatal(err)
  }
  res.Display()
}
