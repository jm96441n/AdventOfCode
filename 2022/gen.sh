#! /bin/zsh

directory="challenge${1}"
mkdir -p $directory

files=(
  "input.txt"
  "test_input.txt"
  "puzzle${1}.go"
  "puzle${1}_test.go"
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
    err, res := $directory.Run(\"./$directory/input.txt\")
    if err != nil {
        log.Fatal(err)
    }
    res.Display()
}" > ./main.go


echo "package $directory

type Result struct {}

func (r Result) Display() {
    fmt.Println(\"done!\")
}

func Run(filename string) Result {

}
" > "$directory/puzzle${1}.go"


echo "package ${directory}_test

import (
	\"testing\"
)

func TestRunPartOne(t *testing.T) {
    res, err := $directory.Run(\"./test_input.txt\")
}
func TestRunPartTwo(t *testing.T) {
    res, err := $directory.Run(\"./test_input.txt\")
}" > "./$directory/puzzle${1}_test.go"


echo "package $directory

func Run() {

}
" > "$directory/puzzle${1}.go"
