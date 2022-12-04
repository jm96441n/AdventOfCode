package main

import (
	"AdventOfCode/2022/challengefour"
	"log"
)

func main() {
	res, err := challengefour.Run("./challengefour/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	res.Display()
}
