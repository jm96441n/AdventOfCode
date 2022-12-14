package main

import (
	"AdventOfCode/2022/challengeseven"
	"fmt"
	"log"
	"os"
)

func main() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	res, err := challengeseven.Run(fmt.Sprintf("%s/challengeseven/input.txt", curDir))
	if err != nil {
		log.Fatal(err)
	}
	res.Display()
}
