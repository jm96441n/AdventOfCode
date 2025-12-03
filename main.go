package main

import (
	"AdventOfCode/2025/challengethree"
  "flag"
  "log"
)

func main() {
	var isTest bool
	flag.BoolVar(&isTest, "test", false, "run test input")
	flag.Parse()
	inputFile := "./2025/challengethree/input.txt"
	if isTest {
		inputFile = "./2025/challengethree/test_input.txt"
	}
  res, err := challengethree.Run(inputFile)
  if err != nil {
    log.Fatal(err)
  }
  res.Display()
}
