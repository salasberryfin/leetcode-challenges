package main

func checkString(s string) bool {
	bFound := false
	for _, i := range s {
		if i != 'a' {
			bFound = true
		} else if bFound {
			return false
		}
	}

	return true
}
