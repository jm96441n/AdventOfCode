package main

import (
	"AdventOfCode/2024/challengefive"
  "log"
)

func main() {
  res, err := challengefive.Run("./2024/challengefive/input.txt")
  if err != nil {
    log.Fatal(err)
  }
  res.Display()
}
