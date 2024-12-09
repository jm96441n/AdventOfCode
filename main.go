package main

import (
	"AdventOfCode/2024/challengetwo"
  "log"
)

func main() {
    res, err := challengetwo.Run("./2024/challengetwo/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    res.Display()
}
