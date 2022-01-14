package main

import "fmt"

// minimize number of calls to isBadVersion
func isBadVersion(version int) bool {
	firstBad := 1
	fmt.Printf("call isBadVersion(%v) --> %v\n", version, version == firstBad)
	return version == firstBad
}

func binarySearch(l, h int, version *int) {
	mid := (l + h) / 2
	if h >= l {
		if isBadVersion(mid) {
			binarySearch(l, mid-1, &mid)
			if mid < *version {
				*version = mid
			}
		} else {
			binarySearch(mid+1, h, version)
		}
	}
}

func firstBadVersion(n int) int {
	firstBad := &n
	binarySearch(1, n, firstBad)
	fmt.Println("First bad version: ", *firstBad)

	return *firstBad
}

func main() {
	input := 1
	firstBadVersion(input)
}
