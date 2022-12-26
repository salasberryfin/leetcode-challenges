package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTreeHeight(tree *TreeNode) int {
	var lHeight, rHeight int

	if tree.Left != nil {
		lHeight += getTreeHeight(tree.Left)
	}
	if tree.Right != nil {
		rHeight += getTreeHeight(tree.Right)
	}

	if lHeight > rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}

func createStr(tree *TreeNode, r, c, height int, result [][]string) {
	x := r + 1
	yLeft := c - int(math.Pow(2, float64(height-r-1)))
	yRight := c + int(math.Pow(2, float64(height-r-1)))
	if tree.Left != nil {
		result[x][yLeft] = fmt.Sprint(tree.Left.Val)
		createStr(tree.Left, x, yLeft, height, result)
	}
	if tree.Right != nil {
		result[x][yRight] = fmt.Sprint(tree.Right.Val)
		createStr(tree.Right, x, yRight, height, result)
	}
}

func printTree(root *TreeNode) [][]string {
	var res [][]string

	res = [][]string{}

	// res is m x n matrix
	// m = height + 1
	// root node position = res[0][(n-1)/2]
	var height int
	height = getTreeHeight(root) - 1
	m := height + 1
	n := int(math.Pow(2, float64(height+1)) - 1)

	rootPos := (n - 1) / 2

	fmt.Printf("Heigth of the tree: %d\n", height)
	fmt.Printf("m=%d, n=%d\n", m, n)
	fmt.Printf("The root node should be in the middle of the top row: (0, %d)\n", rootPos)

	rootLine := make([]string, n)
	for i := 0; i < n; i++ {
		rootLine = append(rootLine, "")
	}

	for i := 0; i < m; i++ {
		line := make([]string, n)
		copy(line, rootLine)
		res = append(res, line)
	}
	res[0][rootPos] = fmt.Sprint(root.Val)

	createStr(root, 0, rootPos, height, res)

	fmt.Printf("Result: %v\n", res)

	return res
}
