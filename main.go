package main

import (
	"AdventOfCode/2024/challengeone"
  "log"
)

func main() {
    res, err := challengeone.Run("./2024/challengeone/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    res.Display()
}
