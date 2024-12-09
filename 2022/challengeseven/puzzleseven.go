package challengeseven

import (
	"AdventOfCode/2022/file"
	"fmt"
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
	name     string
	isDir    bool
	children []*FileNode
	parent   *FileNode
	size     int
}

func (f *FileNode) displayTree(depth int) {
	output := ""
	for i := 0; i < depth; i++ {
		output += "\t"
	}
	typeSym := "F"
	if f.isDir {
		typeSym = "D"
	}
	fmt.Printf("%s%s(%s): %d\n", output, f.name, typeSym, f.size)
	for _, c := range f.children {
		c.displayTree(depth + 1)
	}
}

func Run(filename string) (Result, error) {
	instructions := file.OpenFileIntoStringSlice(filename, file.StringConv)
	pOne, pTwo := parseInstructions(instructions)

	return Result{PartOne: pOne, PartTwo: pTwo}, nil
}

func parseInstructions(instructions []string) (int, int) {
	dirs := make([]*FileNode, 0)
	root := &FileNode{isDir: true, children: make([]*FileNode, 0), name: "/"}
	dirs = append(dirs, root)
	curNode := root
	var name string
	sum := 0
	for i := 0; i < len(instructions); i++ {
		inst := instructions[i]
		switch {
		case inst == "$ cd /":
			curNode = root
		case inst == "$ cd ..":
			curNode = curNode.parent
		case strings.HasPrefix(inst, "$ cd"):
			fmt.Sscanf(inst, "$ cd %s", &name)
			for _, child := range curNode.children {
				if child.name == name {
					curNode = child
				}
			}
		case inst == "$ ls":
			i = parseLSOutput(instructions, i, curNode, &dirs)
		}
	}

	var dfs func(*FileNode) int

	dfs = func(node *FileNode) int {
		if !node.isDir {
			return node.size
		}
		for _, child := range node.children {
			node.size += dfs(child)
		}
		return node.size
	}
	dfs(root)
	existingFreeSpace := 70000000 - dirs[0].size
	deletedSize := dirs[0].size
	neededSpace := 30000000 - existingFreeSpace
	for _, dir := range dirs[1:] {
		if dir.size < 100_000 {
			sum += dir.size
		}
		if dir.size >= neededSpace && dir.size <= deletedSize {
			deletedSize = dir.size
		}
	}
	return sum, deletedSize
}

func parseLSOutput(instructions []string, idx int, curNode *FileNode, dirs *[]*FileNode) int {
	var name string
	var size int
	for {
		idx += 1
		if idx >= len(instructions) {
			return idx
		}
		inst := instructions[idx]
		switch {
		case strings.HasPrefix(inst, "$"):
			idx -= 1
			return idx
		case strings.HasPrefix(inst, "dir "):
			fmt.Sscanf(inst, "dir %s", &name)
			childDir := &FileNode{
				isDir:    true,
				children: make([]*FileNode, 0),
				name:     name,
				parent:   curNode,
			}
			*dirs = append(*dirs, childDir)
			curNode.children = append(curNode.children, childDir)
		default:
			fmt.Sscanf(inst, "%d %s", &size, &name)
			child := &FileNode{
				isDir:  false,
				size:   size,
				name:   name,
				parent: curNode,
			}
			curNode.children = append(curNode.children, child)
		}

	}
}
