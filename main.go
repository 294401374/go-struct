package main

import (
	l "./list"
	"fmt"
)

func main() {
	node := new(l.Node)
	node.Init(5)
	fmt.Println(node)

}
