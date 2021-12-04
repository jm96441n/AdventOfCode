package challenge1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const year int = 2020

func twoProductFromTwoSum(inputs []int, matcher int) (int, int, int) {

	addendOne, addendTwo := TwoSum(inputs, matcher)
	product := addendOne * addendTwo

	return product, addendOne, addendTwo
}

func TwoSum(inputs []int, matcher int) (int, int) {
	diffMap := make(map[int]bool)
	var (
		diff  int
		other int
	)
	for _, rowNum := range inputs {
		diff = matcher - rowNum
		if _, ok := diffMap[diff]; ok {
			other = rowNum
			break
		} else {
			diffMap[rowNum] = true
		}
	}
	return diff, other
}

func threeProductFromThreeSum(inputs []int) int {
	var (
		diff    int
		product int
	)
	for _, rowNum := range inputs {
		diff = year - rowNum
		twoProd, one, two := twoProductFromTwoSum(inputs, diff)
		if twoProd != 0 {
			product = rowNum * one * two
		}
	}
	return product
}

func Run() {
	file, err := os.Open("./challenge1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNums := make([]int, 0)
	for scanner.Scan() {
		rowNum, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		rowNums = append(rowNums, rowNum)
	}
	twoProduct, _, _ := twoProductFromTwoSum(rowNums, year)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Two Product: %d\n", twoProduct)

	threeProduct := threeProductFromThreeSum(rowNums)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Three Product: %d\n", threeProduct)
}
