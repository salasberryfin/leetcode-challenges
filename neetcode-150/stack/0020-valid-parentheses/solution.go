package main

func isValid(s string) bool {
	closeChars := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for _, i := range s {
		if _, ok := closeChars[byte(i)]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != closeChars[byte(i)] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(i))
		}
	}

	return len(stack) == 0
}
