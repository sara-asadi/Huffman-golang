package main

import "fmt"

type Node struct {
	val         string
	freq        int
	left, right *Node
}

func Add(left, right *Node) *Node {
	var newNode Node

	newNode.left = left
	newNode.right = right

	newNode.val = left.val + right.val
	newNode.freq = left.freq + right.freq

	return &newNode
}

func NewNode(val string, freq int) *Node {
	var node = Node{val, freq, nil, nil}
	return &node
}

func (n *Node) Print() {

	if n.left == nil && n.right == nil {
		fmt.Println(n.val)
		return
	} else {
		n.left.Print()
		fmt.Println(n.val)
		n.right.Print()
	}
}

func (n *Node) Encode(code string, encodingMap map[string]string) {

	if n.left == nil && n.right == nil {
		encodingMap[n.val] = code
		return
	} else {
		n.left.Encode(code+"0", encodingMap)
		n.right.Encode(code+"1", encodingMap)
	}
}
