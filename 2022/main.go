package main

import (
	"AdventOfCode/2022/challengeeight"
	"log"
)

func main() {
	res, err := challengeeight.Run("./challengeeight/test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	res.Display()
}
