package challenge7

import (
	"AdventOfCode/utils"
	"fmt"
	"strings"
)

var goldBag string = "shiny gold"

func Run() {
	rows := utils.OpenFileIntoStringSlice("./challenge7/input.txt")
	colorMap := make(map[string]*node)
	trees := buildTrees(rows, colorMap)
	countShinys := 0
	goldTreeNode := colorMap[goldBag]
	for _, tree := range trees {
		if tree == goldTreeNode {
			continue
		}
		if tree.hasShinyBag() {
			countShinys++
		}
	}
	goldBagContainsCount := goldTreeNode.countAllChildren(0)
	fmt.Println(countShinys)
	fmt.Println(goldBagContainsCount)
}

func buildTrees(rows []string, colorMap colorToPointerMap) []*node {
	var (
		identifier    string
		color         string
		topBagColor   string
		childBagColor string
		count         int
		topBag        *node
		ok            bool
		childBag      *node
	)
	topLevelNodes := make([]*node, 0)
	for _, row := range rows {
		rowSplit := strings.Split(row, " contain ")
		fmt.Sscanf(rowSplit[0], "%s %s bags", &identifier, &color)
		topBagColor = fmt.Sprintf("%s %s", identifier, color)

		childrenSplit := strings.Split(rowSplit[1], ", ")
		children := make(map[*node]int)

		// handle the children
		for _, child := range childrenSplit {
			if strings.Contains(child, "no other bags") {
				continue
			}
			fmt.Sscanf(child, "%d %s %s bag", &count, &identifier, &color)
			childBagColor = fmt.Sprintf("%s %s", identifier, color)
			if childBag, ok = colorMap[childBagColor]; !ok {
				childBag = &node{color: childBagColor}
				colorMap[childBagColor] = childBag
			}

			children[childBag] = count
		}

		// handle top bag
		if topBag, ok = colorMap[topBagColor]; !ok {
			topBag = &node{color: topBagColor}
			colorMap[topBagColor] = topBag
		}
		topBag.children = children
		topLevelNodes = append(topLevelNodes, topBag)
	}
	return topLevelNodes
}

type colorToPointerMap map[string]*node

type node struct {
	children map[*node]int
	color    string
}

func (n *node) hasShinyBag() bool {
	if n.color == goldBag {
		return true
	}

	found := false
	for child := range n.children {
		if child.hasShinyBag() {
			found = true
			break
		}
	}
	return found
}

func (n *node) countAllChildren(childAmt int) int {
	if len(n.children) == 0 {
		return 0
	}
	count := 0
	for child, childCount := range n.children {
		count += childCount + (childCount * child.countAllChildren(childCount))
	}
	return count
}
