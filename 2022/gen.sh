#! /bin/zsh

directory="challenge${1}"
mkdir -p $directory

files=(
  "input.txt"
  "test_input.txt"
  "puzzle${1}.go"
)

for file in "${files[@]}"
do
  touch "${directory}/${file}"
done

echo "package main

import (
	\"AdventOfCode/2022/$directory\"
)

func main() {
    $directory.Run()
}" > ./main.go


echo "package $directory

func Run() {

}
" > "$directory/puzzle${1}.go"
