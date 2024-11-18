package main

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))

	if k == 0 {
		return result
	}

	for i := 0; i < len(code); i++ {
		var val int
		if k > 0 {
			for j := 1; j <= k; j++ {
				val += code[(i+j)%len(code)]
			}
		} else {
			for j := 1; j <= -k; j++ {
				val += code[(i-j+len(code))%len(code)]
			}
		}
		result[i] = val
	}

	return result
}
