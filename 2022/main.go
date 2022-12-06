package main

import (
	"AdventOfCode/2022/challengefive"
	"log"
)

func main() {
	res, err := challengefive.Run("./challengefive/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	res.Display()
}
