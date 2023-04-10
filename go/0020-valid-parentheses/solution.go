package main

func isValid(s string) bool {
	closing := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	var stack []byte
	for _, i := range s {
		if v, ok := closing[byte(i)]; ok {
			if len(stack) == 0 || v != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(i))
		}
	}

	return len(stack) == 0
}
