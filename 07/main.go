package main

import "math"

func main() {

}

type MyListNode struct {
	Val  int
	Next *MyListNode
	Prev *MyListNode
}

func (l *MyListNode) toNumber() int {
	n := 0
	node := l
	i := 0
	for node != nil {
		n = n + (node.Val * int(math.Pow(10, float64(i))))

		node = node.Prev
		i++
	}
	return n
}

func reverse(x int) int {
	if x > -9 && x < 9 {
		return x
	}
	neg := false
	if x < 0 {
		neg = true
	}

	l1 := numberToListNode(x)
	x = l1.toNumber()

	if x > math.MaxInt32 || x < -math.MaxUint32 {
		return 0
	} else if !neg {
		return x
	}
	return x * -1
}

func numberToListNode(x int) *MyListNode {
	head := &MyListNode{}
	current := head
	x = int(math.Abs(float64(x)))
	for x > 0 {
		next := &MyListNode{Prev: current}

		n := x % 10

		current.Val = n
		x = x / 10
		if x == 0 {
			break
		}
		current.Next = next
		current = next
	}
	return current
}
