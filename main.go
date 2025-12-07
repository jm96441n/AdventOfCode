package main

import (
	"AdventOfCode/2025/challengesix"
  "flag"
  "log"
)

func main() {
	var isTest bool
	flag.BoolVar(&isTest, "test", false, "run test input")
	flag.Parse()
	inputFile := "./2025/challengesix/input.txt"
	if isTest {
		inputFile = "./2025/challengesix/test_input.txt"
	}
  res, err := challengesix.Run(inputFile)
  if err != nil {
    log.Fatal(err)
  }
  res.Display()
}
