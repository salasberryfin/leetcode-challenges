package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) convert() string {
	var sum string

	sum = fmt.Sprint(l.Val) + sum
	if l.Next != nil {
		sum = l.Next.convert() + sum
	}

	return sum
}

func reverseString(str string) []byte {
	result := make([]byte, len(str))
	i, j := 0, len(str)-1
	for i < j {
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
	strL1, strL2 := l1.convert(), l2.convert()
	intL1, err := strconv.Atoi(strL1)
	intL2, err := strconv.Atoi(strL2)
	if err != nil {
		log.Fatal("Failed to convert string to integer:", err)
	}

	value := intL1 + intL2
	fmt.Printf("The sum of both list is %d\n", value)
	resultString := reverseString(fmt.Sprint(value))

	return createListNode(string(resultString))
}

func printLinkedList(list *ListNode) {
	fmt.Printf("Value: %d\n", list.Val)
	if list.Next != nil {
		printLinkedList(list.Next)
	}
}

func main() {
	example1a := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	example1b := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	example2a := ListNode{
		Val:  0,
		Next: nil,
	}
	example2b := ListNode{
		Val:  0,
		Next: nil,
	}
	example3a := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	example3b := ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
	}

	var start time.Time
	start = time.Now()
	printLinkedList(addTwoNumbers(&example1a, &example1b))
	fmt.Printf("%v\n", time.Since(start))
	fmt.Println()
	start = time.Now()
	printLinkedList(addTwoNumbers(&example2a, &example2b))
	fmt.Printf("%v\n", time.Since(start))
	fmt.Println()
	start = time.Now()
	printLinkedList(addTwoNumbers(&example3a, &example3b))
	fmt.Printf("%v\n", time.Since(start))
}
