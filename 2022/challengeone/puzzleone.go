package challengeone

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

func partOne() (int, int) {
	fileName := "./challengeone/input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	max := 0
	sum := 0
	hp := &IntHeap{}
	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {

			if sum > max {
				max = sum
			}
			heap.Push(hp, -sum)
			sum = 0
			continue

		}
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += i
	}
	heap.Push(hp, -sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	heap.Init(hp)
	totalSum := 0
	for i := 0; i < 3; i++ {
		j := heap.Pop(hp)
		totalSum += j.(int)
	}
	return max, totalSum * -1
}

func Run() {
	one, two := partOne()
	fmt.Printf("Part One: %d\nPart Two %d\n", one, two)
}

type IntHeap []int

func (i *IntHeap) Len() int           { return len(*i) }
func (i *IntHeap) Less(n, j int) bool { return (*i)[n] < (*i)[j] }
func (i *IntHeap) Swap(n, j int)      { (*i)[n], (*i)[j] = (*i)[j], (*i)[n] }

func (i *IntHeap) Push(x any) {
	*i = append(*i, (x).(int))
}

func (i *IntHeap) Pop() any {
	old := *i
	n := len(old)
	x := old[n-1]
	*i = old[0 : n-1]
	return x
}
