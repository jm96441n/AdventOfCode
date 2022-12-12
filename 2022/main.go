package main

import (
	"AdventOfCode/2022/challengeseven"
        "log"
)

func main() {
    res, err := challengeseven.Run("./challengeseven/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    res.Display()
}
