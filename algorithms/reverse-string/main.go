package main

func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		temp := s[l]
		s[l] = s[r]
		s[r] = temp
		l++
		r--
	}
	//fmt.Println("Result:", string(s))
}

func main() {
	inputs := [][]byte{
		[]byte("hello"),
		[]byte("hannah"),
	}

	var x int
	for x < len(inputs) {
		reverseString(inputs[x])
		x++
	}
}
