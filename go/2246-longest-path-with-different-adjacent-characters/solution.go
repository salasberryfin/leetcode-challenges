/*
Just can't wrap my head around this one...
Had to review the official LeetCode solution and yet I do not fully understand
why it keeps the two longest paths.
Hope to come back to this one in the future and be able to do it by myself

Keys
	- Use dfs
	- Keep track of the 2 longest paths for each node: max1, max2
	- For each iteration, update the values of these two paths
	- Update the solution value based on the sum of these two paths
	- Count the current node as well (+1)
*/
package main

func dfs(routes map[int][]int, node, parent int, chars string, longest *int) int {
	visited := map[int]bool{}
	visited[node] = true
	max1, max2 := 0, 0

	for _, w := range routes[node] {
		if parent != w {
			if _, ok := visited[w]; !ok {
				childLength := dfs(routes, w, node, chars, longest)

				if chars[w] != chars[node] {
					if childLength > max1 {
						max2 = max1
						max1 = childLength
					} else if childLength > max2 {
						max2 = childLength
					}
				}
			}
		}
	}
	if *longest < max1+max2+1 {
		*longest = max1 + max2 + 1
	}

	return max1 + 1
}

func longestPath(parent []int, s string) int {
	routes := map[int][]int{}
	for i := 1; i < len(parent); i++ {
		routes[parent[i]] = append(routes[parent[i]], i)
		routes[i] = append(routes[i], parent[i])
	}

	longest := 1
	dfs(routes, 0, -1, s, &longest)

	return longest
}
