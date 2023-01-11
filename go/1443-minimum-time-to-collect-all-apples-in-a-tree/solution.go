/*
 * Input
 * - n: vertices
 * - edges: meaning edges[i] = [ai, bi] -> connection between ai and bi
 * - hasApple: hasApple[i] -> vertex i has an apple
 *
 * Went through hell with this one
 * Even implemented dijkstra to try to find shortest paths between apples -> not a good idea
 *
 * In the end, with some help from internet, the plan is to:
 *  - We first use a hash map to store the neighbors of each node in the graph: store neighbors both ways!
 *	- Use a Depth first search (DFS) to traverse the graph (it is a grap, not a binary tree)
 *	- Implement it recursively so that we can detect which branch has apples
 *	- To avoid going over the same pair of parent/child nodes: check for root != parent
 *	- If we detect an apple, we have to sum +2: count getting the apple and coming back to the parent
 *  - If we get a positive result from the recursive call: add it to the result variable
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

func dfs(routes map[int][]int, hasApples []bool, root, parent int) int {
	var result int

	var stack []int
	var visited []int

	stack = append(stack, root)

	visited = append(visited, root)
	for _, adjacent := range routes[root] {
		if adjacent != parent {
			time := dfs(routes, hasApples, adjacent, root)
			if time > 0 || hasApples[adjacent] {
				result += 2 + time
			}
		}
	}

	return result
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	var apples []int

	for i, apple := range hasApple {
		if apple {
			apples = append(apples, i)
		}
	}

	if len(apples) == 0 {
		return 0
	}

	var routes = map[int][]int{}
	for _, edge := range edges {
		routes[edge[0]] = append(routes[edge[0]], edge[1])
		routes[edge[1]] = append(routes[edge[1]], edge[0])
	}

	return dfs(routes, hasApple, 0, -1)
}
