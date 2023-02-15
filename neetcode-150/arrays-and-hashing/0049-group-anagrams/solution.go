/*
1. Iterate over each element
2. Merge Sort
3. Check if sorted element in hash map
4. Store sorted in hash map: map[sorted]=[]string{original}
*/
package main

func sort(l, r string) string {
	var result string
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] > r[j] {
			result += string(r[j])
			j++
		} else {
			result += string(l[i])
			i++
		}
	}

	for i < len(l) {
		result += string(l[i])
		i++
	}
	for j < len(r) {
		result += string(r[j])
		j++
	}

	return result
}

func merge(strs string) string {
	if len(strs) <= 1 {
		return strs
	}

	mid := len(strs) / 2
	left := merge(strs[:mid])
	right := merge(strs[mid:])

	return sort(left, right)
}

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}
	for _, str := range strs {
		sorted := merge(str)
		hashMap[sorted] = append(hashMap[sorted], str)
	}
	var result [][]string
	for _, v := range hashMap {
		result = append(result, v)
	}

	return result
}
