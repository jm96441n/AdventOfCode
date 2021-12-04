package main

import (
	"AdventOfCode/2020/challenge1"
	"AdventOfCode/2020/challenge10"
	"AdventOfCode/2020/challenge11"
	"AdventOfCode/2020/challenge12"
	"AdventOfCode/2020/challenge13"
	"AdventOfCode/2020/challenge14"
	"AdventOfCode/2020/challenge15"
	"AdventOfCode/2020/challenge16"
	"AdventOfCode/2020/challenge17"
	"AdventOfCode/2020/challenge18"
	"AdventOfCode/2020/challenge19"
	"AdventOfCode/2020/challenge2"
	"AdventOfCode/2020/challenge3"
	"AdventOfCode/2020/challenge4"
	"AdventOfCode/2020/challenge5"
	"AdventOfCode/2020/challenge7"
	"AdventOfCode/2020/challenge8"
	"AdventOfCode/2020/challenge9"
	"os"
)

func main() {
	challengeToRun := os.Args[1]
	switch challengeToRun {
	case "1":
		challenge1.Run()
	case "2":
		challenge2.Run()
	case "3":
		challenge3.Run()
	case "4":
		challenge4.Run()
	case "5":
		challenge5.Run()
	case "7":
		challenge7.Run()
	case "8":
		challenge8.Run()
	case "9":
		challenge9.Run()
	case "10":
		challenge10.Run()
	case "11":
		challenge11.Run()
	case "12":
		challenge12.Run()
	case "13":
		challenge13.Run()
	case "14":
		challenge14.Run()
	case "15":
		challenge15.Run()
	case "16":
		challenge16.Run()
	case "17":
		challenge17.Run()
	case "18":
		challenge18.Run()
        case "19":
		challenge19.Run()
	}

}
