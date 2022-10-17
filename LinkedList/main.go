package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

func (l linkedList) len() int {
	return l.length
}


func (l linkedList) display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.next
	}
}

func (l *linkedList) PushBack(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func main() {
	link := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 10}
	node3 := &node{data: 10}
	node4 := &node{data: 10}
	link.PushBack(node1)
	link.PushBack(node2)
	link.PushBack(node3)
	link.PushBack(node4)
	link.display()
	link.len()
}
