package challengeseven

import (
	"AdventOfCode/2022/file"
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	PartOne int
	PartTwo int
}

func (r Result) Display() {
	fmt.Printf("Part One: %d\nPart Two: %d\n", r.PartOne, r.PartTwo)
}

type FileNode struct {
	isDir    bool
	children []*FileNode
	parent   *FileNode
	size     int
}

func (f *FileNode) displayTree() {
	fmt.Println("hello")
}

func Run(filename string) (Result, error) {
	instructions := file.OpenFileIntoSlice(filename, file.StringConv)
	fileTree := parseInstructions(instructions)
	fileTree.displayTree()

	return Result{}, nil
}

func parseInstructions(instructions []string) *FileNode {
	nodes := make(map[string]*FileNode, 0)
	root := &FileNode{isDir: true, children: make([]*FileNode, 0)}
	nodes["/"] = root
	curNode := root
	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]
		fmt.Println(inst)
		switch {
		case inst == "$ cd /":
			curNode = root
		case inst == "$ cd ..":
			curNode = curNode.parent
		case strings.HasPrefix(inst, "$ cd"):
			val := strings.TrimPrefix(inst, "$ cd ")
			fmt.Println(nodes)
			curNode = nodes[val]
		case inst == "$ ls":
			i = parseLSOutput(instructions, i, curNode, nodes)
		}
	}
	return root
}

func parseLSOutput(instructions []string, idx int, curNode *FileNode, nodes map[string]*FileNode) int {
	for {
		idx += 1
		inst := instructions[idx]
		fmt.Println(inst)
		switch {
		case strings.HasPrefix(inst, "$"):
			idx -= 1
			return idx
		case strings.HasPrefix(inst, "dir "):
			name := strings.TrimPrefix("dir ", inst)
			childDir := &FileNode{isDir: true, children: make([]*FileNode, 0)}
			nodes[name] = childDir
			fmt.Println(curNode)
			curNode.children = append(curNode.children, childDir)
		default:
			vals := strings.Split(inst, " ")
			name := vals[1]
			size, err := strconv.Atoi(vals[0])
			if err != nil {
				panic(err)
			}
			child := &FileNode{isDir: false, children: make([]*FileNode, 0), size: size}
			nodes[name] = child
			curNode.children = append(curNode.children, child)
		}

	}
}
