package main

func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return 1
	}
	personTrusts := map[int][]int{}
	personIsTrustedBy := map[int][]int{}
	maxTrusts := 0
	judge := -1
	for _, i := range trust {
		personTrusts[i[0]] = append(personTrusts[i[0]], i[1])
		personIsTrustedBy[i[1]] = append(personIsTrustedBy[i[1]], i[0])
		if len(personIsTrustedBy[i[1]]) > maxTrusts {
			maxTrusts = len(personIsTrustedBy[i[1]])
			judge = i[1]
		}
	}
	if maxTrusts == n-1 && len(personTrusts[judge]) == 0 {
		return judge
	}

	return -1
}
