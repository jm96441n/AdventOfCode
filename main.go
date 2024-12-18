package main

import (
	"AdventOfCode/2024/challengethree"
	"log"
)

func main() {
	res, err := challengethree.Run("./2024/challengethree/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	res.Display()
}
