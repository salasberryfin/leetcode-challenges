package main

func smallestUniqueSubarray(nums []int) int {
	result := len(nums)

	// set left and right boundaries
	left, right := 1, len(nums)
	for left <= right {
		mid := (left + right) / 2
		if hasUniqueSubarray(nums, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func hasUniqueSubarray(nums []int, mid int) bool {
	const (
		base1 int64 = 100003
		mod1  int64 = 1000000007
		base2 int64 = 100019
		mod2  int64 = 1000000009
	)

	var h1, h2 int64
	var pow1, pow2 int64 = 1, 1

	// Build the initial window hash and precompute base^(mid-1) % mod
	for i := range mid {
		val := int64(nums[i]) + 1 // +1 avoids zero-value collisions
		h1 = (h1*base1 + val) % mod1
		h2 = (h2*base2 + val) % mod2

		if i < mid-1 {
			pow1 = (pow1 * base1) % mod1
			pow2 = (pow2 * base2) % mod2
		}
	}

	counts := make(map[[2]int64]int)
	counts[[2]int64{h1, h2}]++

	// Slide the window across the rest of the array
	for i := mid; i < len(nums); i++ {
		outVal := int64(nums[i-mid]) + 1
		inVal := int64(nums[i]) + 1

		// Remove outgoing element and add incoming element
		h1 = (h1 - (outVal*pow1)%mod1 + mod1) % mod1
		h1 = (h1*base1 + inVal) % mod1

		h2 = (h2 - (outVal*pow2)%mod2 + mod2) % mod2
		h2 = (h2*base2 + inVal) % mod2

		counts[[2]int64{h1, h2}]++
	}

	// Check if any subarray hash occurred only once
	for _, count := range counts {
		if count == 1 {
			return true
		}
	}

	return false
}
