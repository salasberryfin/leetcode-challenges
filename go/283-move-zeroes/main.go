package main

import "fmt"

func moveZeroesInefficient(nums []int) {
	for i := 0; i < (len(nums) - 1); i++ {
		for j := 0; j < (len(nums) - 1); j++ {
			if nums[j] == 0 {
				nums[j] = nums[j+1]
				nums[j+1] = 0
			}
		}
	}

	fmt.Println("Result for", nums, "-", nums)
}

func split(nums []int, l, r int) {

}

func moveZeroes(nums []int) {
	fmt.Println("Current nums:", nums)
	var i int
	var zeroes int
	for zeroes+i < len(nums)-1 {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			zeroes++
		} else {
			i++
		}
		fmt.Println("After iteration", i+1, "nums is", nums)
	}

	fmt.Println("Result:", nums)
}

func main() {
	inputs := [][]int{
		{0, 1, 0, 3, 12},
		{0},
		{0, 0, 1},
	}
	for _, input := range inputs {
		moveZeroes(input)
		//moveZeroesInefficient(input)
	}
}
