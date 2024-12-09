#! /bin/zsh

directory="challenge${1}"
mkdir -p $directory

files=(
  "input.txt"
  "test_input.txt"
  "puzzle${1}.go"
  "puzzle${1}_test.go"
)

for file in "${files[@]}"
do
  touch "${directory}/${file}"
done

echo "package main

import (
	\"AdventOfCode/2022/$directory\"
        \"log\"
)

func main() {
    res, err := $directory.Run(\"./$directory/input.txt\")
    if err != nil {
        log.Fatal(err)
    }
    res.Display()
}" > ./main.go


echo "package $directory

type Result struct {
    PartOne int
    PartTwo int
}

func (r Result) Display() {
    fmt.Printf(\"PartOne: %d\nPartTwo: %d\n\", r.PartOne, r.PartTwo)
}

func Run(filename string) (Result, error) {
    return Result{}, nil
}
" > "./$directory/puzzle${1}.go"


echo "package ${directory}_test

import (

	\"AdventOfCode/2022/${directory}\"
	\"testing\"
)

func TestRunPartOne(t *testing.T) {
        res, err := $directory.Run(\"./test_input.txt\")
	if err != nil {
		t.Error(err)
	}
	if res.PartOne != 0 {
		t.Errorf(\"Expected: 0, Got: %d\", res.PartOne)
	}

}
func TestRunPartTwo(t *testing.T) {
        res, err := $directory.Run(\"./test_input.txt\")
	if err != nil {
		t.Error(err)
	}
	if res.PartTwo != 0 {
		t.Errorf(\"Expected: 0, Got: %d\", res.PartTwo)
	}

}" > "./$directory/puzzle${1}_test.go"



day=$(($(ls | wc -l) - 3))
curl --cookie $COOKIE https://adventofcode.com/2023/day/$day/input > ./$directory/input.txt
