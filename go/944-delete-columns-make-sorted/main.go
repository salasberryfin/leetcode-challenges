package main

func minDeletionSize(strs []string) int {
	//start := time.Now()
	var counter int
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				counter++
				break
			}
		}
	}
	//fmt.Printf("Elapsed time optimal method: %dms\n", time.Since(start))

	return counter
}
