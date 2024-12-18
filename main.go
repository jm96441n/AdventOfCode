package main

import (
	"AdventOfCode/2024/challengefour"
  "log"
)

func main() {
    res, err := challengefour.Run("./2024/challengefour/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    res.Display()
}
