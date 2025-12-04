package main

import (
	"AdventOfCode/2025/challengefive"
  "flag"
  "log"
)

func main() {
	var isTest bool
	flag.BoolVar(&isTest, "test", false, "run test input")
	flag.Parse()
	inputFile := "./2025/challengefive/input.txt"
	if isTest {
		inputFile = "./2025/challengefive/test_input.txt"
	}
  res, err := challengefive.Run(inputFile)
  if err != nil {
    log.Fatal(err)
  }
  res.Display()
}
