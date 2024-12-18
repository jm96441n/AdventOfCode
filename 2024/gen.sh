#! /bin/zsh

directory="2024/challenge${1}"
packageName="challenge${1}"
mkdir -p "$directory"

files=(
  "input.txt"
  "test_input.txt"
  "puzzle${1}.go"
  "puzzle${1}_test.go"
)

for file in "${files[@]}"; do
  touch "${directory}/${file}"
done

echo "package main

import (
	\"AdventOfCode/2024/$packageName\"
  \"log\"
)

func main() {
    res, err := $packageName.Run(\"./$directory/input.txt\")
    if err != nil {
        log.Fatal(err)
    }
    res.Display()
}" >./main.go

echo -E "package $packageName

import \"fmt\"

type Result struct {
    PartOne int
    PartTwo int
}

func (r Result) Display() {
  fmt.Printf(\"PartOne: %d\\nPartTwo: %d\\n\", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
  return Result{}, nil
}
" >"./$directory/puzzle${1}.go"

echo "package ${packageName}_test

import (

	\"AdventOfCode/2024/${packageName}\"
	\"testing\"
)

func TestRunPartOne(t *testing.T) {
  res, err := $packageName.Run(\"./test_input.txt\")
	if err != nil {
		t.Error(err)
	}

  expected := 0

	if res.PartOne != expected {
		t.Errorf(\"Expected: %d, Got: %d\", expected, res.PartOne)
	}

}
func TestRunPartTwo(t *testing.T) {
  res, err := $packageName.Run(\"./test_input.txt\")
	if err != nil {
		t.Error(err)
	}

  expected := 0

	if res.PartTwo != expected {
		t.Errorf(\"Expected: %d, Got: %d\", expected, res.PartTwo)
	}

}" >"./$directory/puzzle${1}_test.go"

day=$(($(ls ./2024 | wc -l) - 1))
curl --cookie $COOKIE https://adventofcode.com/2024/day/$day/input >./$directory/input.txt
