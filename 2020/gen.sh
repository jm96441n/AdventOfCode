#! /bin/zsh

directory="Challenge${1}"
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
