package main

import "fmt"

func neighbors(row, col, m, n int) [][]int {
	var result [][]int
	// up
	if row+1 < m {
		result = append(result, []int{row + 1, col})
	}
	// down
	if row-1 >= 0 {
		result = append(result, []int{row - 1, col})
	}
	// right
	if col+1 < n {
		result = append(result, []int{row, col + 1})
	}
	// left
	if col-1 >= 0 {
		result = append(result, []int{row, col - 1})
	}

	return result
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m, n := len(image), len(image[0])
	startingColor := image[sr][sc]

	queue := make([][]int, 0)
	queue = append(queue, []int{sr, sc})
	explored := map[int]bool{}
	// breadth-first search
	for len(queue) > 0 {
		node := queue[0]
		//fmt.Println("Current node:", node)
		queue = queue[1:]

		for _, neighbor := range neighbors(node[0], node[1], m, n) {
			// iterate over list of 4-direction neighbors
			//fmt.Println("Neighbor:", neighbor)
			neighborColor := image[neighbor[0]][neighbor[1]]
			_, ok := explored[neighbor[0]+neighbor[1]*m]
			if neighborColor == startingColor && !ok {
				// neighbor has not been explored and color matches
				queue = append(queue, neighbor)
			}

			explored[neighbor[0]+neighbor[1]*m] = true
		}

		//fmt.Println("Queue after neighbor search:", queue)

		// set neighbor as explored
		// update color
		image[node[0]][node[1]] = newColor

		//fmt.Println("Updated result grid:", image)
	}

	return image
}

func main() {
	// image[i][j] pixels of the m x n image
	// floodFill from image[sr][sc]
	inputs := [][][]int{
		{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		},
		{
			{0, 0, 0},
			{0, 0, 0},
		},
	}

	fmt.Println("Result for", inputs[0], "is\n", floodFill(inputs[0], 1, 1, 2))
	fmt.Println("Result for", inputs[1], "is\n", floodFill(inputs[1], 0, 0, 2))

}
