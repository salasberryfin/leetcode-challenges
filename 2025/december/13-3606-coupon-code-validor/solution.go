package main

import (
	"sort"
	"unicode"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	var result []string

	validLinesSlice := []string{"electronics", "grocery", "pharmacy", "restaurant"}
	validLines := map[string][]string{
		"electronics": {},
		"grocery":     {},
		"pharmacy":    {},
		"restaurant":  {},
	}

	for i, v := range code {
		if isActive[i] {
			// coupon is active
			if _, ok := validLines[businessLine[i]]; ok {
				// it is a valid business line
				line := businessLine[i]
				if isAlphanumeric(v) {
					// it is a valid coupon code
					validLines[line] = append(validLines[line], v)
				}
			}
		}
	}

	for _, v := range validLinesSlice {
		// sort alphabetically on `code` before appending to result
		sort.Strings(validLines[v])
		result = append(result, validLines[v]...)
	}

	return result
}

func isAlphanumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}

	return true
}
