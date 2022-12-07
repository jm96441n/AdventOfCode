package main

import (
	"AdventOfCode/2022/challengesix"
	"log"
)

func main() {
	res, err := challengesix.Run("./challengesix/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	res.Display()
}
