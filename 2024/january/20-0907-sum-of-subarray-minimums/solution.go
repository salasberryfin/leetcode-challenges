/*
* This is definitely a hard one: check monotonic stack.
 */
package main

func sumSubarrayMins(arr []int) int {
	var result int
	modulo := 1000000007

	stack := []int{}

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			var left, right int
			currentIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left = currentIndex - stack[len(stack)-1]
			} else {
				left = currentIndex + 1
			}
			right = i - currentIndex
			result = (result + arr[currentIndex]*left*right) % modulo
		}
		stack = append(stack, i)
	}

	for i := range stack {
		var left, right int
		currentIndex := stack[i]
		if i > 0 {
			left = currentIndex - stack[i-1]
		} else {
			left = currentIndex + 1
		}
		right = len(arr) - currentIndex
		result = (result + arr[currentIndex]*left*right) % modulo
	}

	return result
}
