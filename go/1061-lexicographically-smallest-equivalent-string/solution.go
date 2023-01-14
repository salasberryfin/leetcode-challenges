/*
Not a particularly difficult problem but need to understand what type of data structure
is required to solve it.

Keys:
	- It is a graph problem. We can treat each character in a string as nodes related
	to the characters in the other string (s1, s2)
	- Create and ajacency array that stores the equivalencies between s1 and s2
	- We'll use this array to analyze each case and create the union.
		- e.g if a = b and b = c, a = c
	- Using DFS we go over the whole list of adjacent 'nodes'
		- For each, we need to detect the minimum character
		- Set the queue and start analyzing 'children'
		- if the value is less than the current min, we set it as new min
		- we keep track of each of the associated 'nodes'
			- then we'll assign the minimum character to each one of the positions of these 'nodes'
			in the `minEquivalent` array
		- `minEquivalent` is a length 26 array with the minimum character associated with each
		lower case letter of the alphabet
*/
package main

func dfs(adjacent [][]int) []byte {
	var visited = make([]bool, 26)
	var minEquivalent = make([]byte, 26)

	for i := range adjacent {
		if !visited[i] {
			min := 27
			queue := []int{}
			queue = append(queue, i)
			equivalent := []int{}
			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]
				if !visited[current] {
					//fmt.Printf("For current %c: ", current+'a')
					visited[current] = true
					equivalent = append(equivalent, current)
					if current < min {
						min = current
					}
					for _, neighbors := range adjacent[current] {
						queue = append(queue, neighbors)
					}
				}
			}
			for _, set := range equivalent {
				minEquivalent[set] = byte(min + 'a')
				//fmt.Printf("the minimum is %c\n", minEquivalent[set])
			}
		}
	}

	return minEquivalent
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	var result string
	var adjacent = make([][]int, 26)
	for i := range s1 {
		charS1, charS2 := int(s1[i]-'a'), int(s2[i]-'a')
		adjacent[charS1] = append(adjacent[charS1], charS2)
		adjacent[charS2] = append(adjacent[charS2], charS1)
	}
	equivalency := dfs(adjacent)
	for _, i := range baseStr {
		result += string(equivalency[i-'a'])
	}
	//fmt.Printf("Result is %s\n", result)

	return result
}
