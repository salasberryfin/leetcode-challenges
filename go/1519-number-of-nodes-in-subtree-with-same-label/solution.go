/*
Can be solved using Depth-first search (DFS)
Keys:
	- Get the labels for each child when recursively calling DFS
		- This can be done by creating an array of len 26 = # lower case characters
		- Update to 1 the position of the character that corresponds to the label of the node
			e.g. label c -> array['c' - 'a'] -> this converts the hex value to a 0-index
		- Do this for each child and, when calling the recursive function, sum each of the items of the child's
		array to the items of the parent's array
		- Before any operation, check that we're not going parent-child indefinitely
			parent != current
		- The result will be stored in array of length # of nodes
		- Before returning the values for the given node:
			update the position of the current node in the result array
			this will contain its result and its childs
		- When all the nodes are visited and the result is ready
			it can be used from the original function, since we were updating the same memory locations in the dfs function
*/
package main

func contains(slice []int, val int) bool {
	for _, i := range slice {
		if i == val {
			return true
		}
	}

	return false
}

func sum(left, right [26]int) [26]int {
	var result [26]int

	for i := 0; i < len(left); i++ {
		result[i] = left[i] + right[i]
	}

	return result
}

func dfs(routes map[int][]int, labels string, root, parent int, result []int) [26]int {
	visited := []int{}
	visited = append(visited, root)
	values := [26]int{}
	values[labels[root]-'a'] = 1

	for _, w := range routes[root] {
		if w != parent {
			if !contains(visited, w) {
				childCount := dfs(routes, labels, w, root, result)
				for i := 0; i < len(values); i++ {
					values[i] += childCount[i]
				}
			}
		}
	}
	result[root] = values[labels[root]-'a']

	return values
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	routes := map[int][]int{}

	for _, edge := range edges {
		routes[edge[0]] = append(routes[edge[0]], edge[1])
		routes[edge[1]] = append(routes[edge[1]], edge[0])
	}

	result := make([]int, len(labels))
	dfs(routes, labels, 0, -1, result)

	return result
}
