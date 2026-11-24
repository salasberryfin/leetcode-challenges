package main

func maxLength(nums []int) int {
	// prod(arr) == lcm(arr) * gcd(arr)

	result := 1

	maxLCM, maxVal := 1, 0
	for _, x := range nums {
		maxLCM = lcm(maxLCM, x)
		maxVal = max(maxVal, x)
	}
	limit := maxLCM * maxVal

	for i := range nums {
		prod, g, l := 1, 0, 1
		for j := i; j < len(nums); j++ {
			prod = prod * nums[j]
			if prod > limit {
				break
			}
			g = gcd(g, nums[j])
			l = lcm(l, nums[j])
			if prod == g*l {
				result = max(result, j-i+1)
			}
		}
	}

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}

func lcm(a, b int) int {
	return abs(a*b) / gcd(a, b)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
