package challengethree

import "fmt"

type Result struct {
    PartOne int
    PartTwo int
}

func (r Result) Display() {
  fmt.Printf("PartOne: %d\nPartTwo: %d\n", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
  return Result{}, nil
}

