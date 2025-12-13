package main

import (
	"AdventOfCode/2025/challengeeight"
  "flag"
  "log"
)

func main() {
	var isTest bool
	flag.BoolVar(&isTest, "test", false, "run test input")
	flag.Parse()
	inputFile := "./2025/challengeeight/input.txt"
	if isTest {
		inputFile = "./2025/challengeeight/test_input.txt"
	}
  res, err := challengeeight.Run(inputFile)
  if err != nil {
    log.Fatal(err)
  }
  res.Display()
}
