/*
Printing the nodes in Traversal order is trivial.
	1. Recursively iterate over first Left, then Right and print value just
	before recursion
Storing the values, per level, in a matrix is more complex
Can be implemented with a queue
I opted for the pointer solution, that seems easier
Instead of populating an array within the recursive function, which I was not able
to do, I passed the result matrix as a pointer.
When we pass through a element of the level for the first time -len(result)==level-
we need to initialize with a value, so that we can access by index -the level- later
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(root *TreeNode, level int, arr *[][]int) {
	if root == nil {
		return
	}

	if len(*arr) == level {
		*arr = append(*arr, []int{root.Val})
	} else {
		//fmt.Printf("Level %d, Val %d\n", level, root.Val)
		(*arr)[level] = append((*arr)[level], root.Val)
	}

	traversal(root.Left, level+1, arr)
	traversal(root.Right, level+1, arr)
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	traversal(root, 0, &result)
	//fmt.Printf("Content %v \n", result)

	return result
}
