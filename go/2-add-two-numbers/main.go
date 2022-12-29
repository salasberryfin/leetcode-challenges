package main

import (
	"fmt"
	"strconv"
)

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func carrySum(str1, str2 string) string {
	var result string
	var carry int

	var i, j int
	for i < len(str1) && j < len(str2) {
		int1 := int(str1[i]) - int('0')
		int2 := int(str2[i]) - int('0')
		sum := int1 + int2 + carry
		carry = sum / 10
		result = result + fmt.Sprint(sum%10)
		i++
		j++
	}

	for i < len(str1) {
		int1 := int(str1[i]) - int('0')
		sum := int1 + carry
		carry = sum / 10
		result = result + fmt.Sprint(sum%10)
		i++
	}
	for j < len(str2) {
		int2 := int(str2[j]) - int('0')
		sum := int2 + carry
		carry = sum / 10
		result = result + fmt.Sprint(sum%10)
		j++
	}

	if carry != 0 {
		result = result + fmt.Sprint(carry)
	}

	return result
}

func convert(l *ListNode) string {
	var sum string

	sum = fmt.Sprint(l.Val) + sum
	if l.Next != nil {
		sum = convert(l.Next) + sum
	}

	return sum
}

func reverseString(str string) []byte {
	result := make([]byte, len(str))
	i, j := 0, len(str)-1
	for i <= j {
		result[i] = str[j]
		result[j] = str[i]
		i++
		j--
	}

	return result
}

func createListNode(result string) *ListNode {
	var resultList ListNode

	stringV := string(result[0])
	intV, _ := strconv.Atoi(stringV)
	resultList.Val = intV
	if len(result)-1 > 0 {
		resultList.Next = createListNode(result[1:])
	} else {
		resultList.Next = nil
	}

	return &resultList
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	strL1, strL2 := convert(l1), convert(l2)
	op1, op2 := string(reverseString(strL1)), string(reverseString(strL2))

	carryValue := carrySum(op1, op2)
	fmt.Printf("The sum of both lists is %s\n", carryValue)

	return createListNode(carryValue)
}

func printLinkedList(list *ListNode) {
	fmt.Printf("Value: %d\n", list.Val)
	if list.Next != nil {
		printLinkedList(list.Next)
	}
}
