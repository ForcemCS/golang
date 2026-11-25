package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main6() {

	//链表
	node1 := &Node{Value: 10}
	node2 := &Node{Value: 20}
	node3 := &Node{Value: 30}

	node1.Next = node2
	node2.Next = node3

	for current := node1; current != nil; current = current.Next {
		fmt.Println(current.Value)
	}
}
