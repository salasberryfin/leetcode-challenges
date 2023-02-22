package main

func hammingWeight(num uint32) int {
	ones := 0
	for num != 0 {
		// check if the least significant bit is a 1
		ones += int(num & 1) // this can be done with nums AND 1 or with num%2
		// shift bits to the right: discard least significant bit
		num = num >> 1
	}

	return ones
}
