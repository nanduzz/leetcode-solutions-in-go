package main

import (
	"fmt"
	"math/big"
)

func main() {
	n1 := big.NewInt(0)
	n1.SetString("1000000000000000000000000000001", 10)
	l1 := fromNumber(n1)
	l2 := fromNumber(big.NewInt(465))

	numbers := addTwoNumbers(l1, l2)
	fmt.Println(parseToNumber(numbers))
}

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := parseToNumber(l1)
	n2 := parseToNumber(l2)

	n1.Add(n1, n2)
	return fromNumber(n1)
}

func parseToNumber(l *ListNode) *big.Int {
	mult := big.NewInt(0)

	list := l
	n := big.NewInt(0)

	for list != nil {
		power := big.NewInt(0)
		power.Exp(big.NewInt(10), mult, nil)
		n.Add(n, power.Mul(power, big.NewInt(int64(list.Val))))
		mult.Add(mult, big.NewInt(1))
		list = list.Next
	}
	return n
}

func fromNumber(n *big.Int) *ListNode {

	tenAsBigInt := big.NewInt(10)
	var q, r big.Int
	q.DivMod(n, tenAsBigInt, &r)
	head := &ListNode{Val: int(r.Int64()), Next: nil}
	var current = head
	for q.Cmp(big.NewInt(0)) > 0 {
		q.DivMod(&q, tenAsBigInt, &r)
		next := &ListNode{Val: int(r.Int64()), Next: nil}
		current.Next = next
		current = next
	}
	return head
}
