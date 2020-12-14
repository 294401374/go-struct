package main

import (
	l "./list"
	"fmt"
)

func main() {
	node := new(l.Node)
	node.Init(5)
	node.List()
	fmt.Println("--------------------------")
	head := node
	//newHead := node.ReverseNode(head)
	head.List()
	fmt.Println("---------最后一个的指针和data都是nil----------")
	newHead := node.ReverseNode(head)
	newHead.List()
}
