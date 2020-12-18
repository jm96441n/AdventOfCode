package challenge18

import (
	"AdventOfCode/file_utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	rows := file_utils.OpenFileIntoSlice("challenge18/input.txt")
	partOneSum := 0
	partTwoSum := 0
	for _, row := range rows {
		partOneSum += findSumForExpression(row, partOneConditional)
		partTwoSum += findSumForExpression(row, partTwoConditional)
	}
	fmt.Println(partOneSum)
	fmt.Println(partTwoSum)

}

func findSumForExpression(expr string, iterator iteratorConditional) int {
	postfix := buildPostfix(expr, iterator)
	val := executePostfix(postfix)
	return val
}

var (
	openParens  = "("
	closeParens = ")"
	add         = "+"
	mult        = "*"
)

type iteratorConditional func(s *stack) bool

func buildPostfix(expr string, iterator iteratorConditional) []string {
	bits := strings.Split(expr, "")
	stack := &stack{size: 0}
	exp := make([]string, 0)
	for _, op := range bits {
		if op == " " {
			continue
		}
		switch op {
		case openParens:
			stack.push(op)
		case closeParens:
			{
				poppedNode := stack.pop()
				for poppedNode.val != openParens {
					exp = append(exp, poppedNode.val)
					poppedNode = stack.pop()
				}
			}
		case add:
			{
				for iterator(stack) {
					exp = append(exp, stack.pop().val)
				}
				stack.push(op)

			}
		case mult:
			{
				for multIterator(stack) {
					exp = append(exp, stack.pop().val)
				}

				stack.push(op)

			}

		default:
			exp = append(exp, op)
		}

	}
	for stack.size > 0 {
		exp = append(exp, stack.pop().val)
	}
	return exp
}

func partOneConditional(s *stack) bool {
	return (s.size > 0) && (s.peek() == mult || s.peek() == add)
}

func partTwoConditional(s *stack) bool {
	return (s.size > 0) && s.peek() == add
}

func multIterator(s *stack) bool {
	return (s.size > 0) && (s.peek() == mult || s.peek() == add)
}

func executePostfix(pf []string) int {
	stack := &stack{size: 0}
	for _, op := range pf {
		switch op {
		case add:
			{
				addendOne, err := strconv.Atoi(stack.pop().val)
				if err != nil {
					panic(err)
				}
				addendTwo, err := strconv.Atoi(stack.pop().val)
				if err != nil {
					panic(err)
				}
				sum := addendOne + addendTwo
				stack.push(fmt.Sprintf("%d", sum))

			}
		case mult:
			{
				multOne, err := strconv.Atoi(stack.pop().val)
				if err != nil {
					panic(err)
				}
				multTwo, err := strconv.Atoi(stack.pop().val)
				if err != nil {
					panic(err)
				}
				prod := multOne * multTwo
				stack.push(fmt.Sprintf("%d", prod))

			}
		default:
			stack.push(op)
		}
	}
	if stack.size != 1 {
		panic("Stack should be of length 1")
	}
	topVal, _ := strconv.Atoi(stack.peek())
	return topVal
}

type stack struct {
	head *node
	size int
}

type node struct {
	next *node
	val  string
}

func (s *stack) pop() *node {
	if s.size <= 0 {
		return nil
	}
	newHead := s.head.next
	oldHead := s.head
	s.head = newHead
	s.size--
	return oldHead
}

func (s *stack) push(val string) {
	newHead := &node{
		next: s.head,
		val:  val,
	}
	s.head = newHead
	s.size++

}

func (s *stack) peek() string {
	return s.head.val
}
