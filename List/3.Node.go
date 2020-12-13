package List

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Next *Node
	Data interface{}
}

// 获取第i个元素
func (node *Node) Get(i int) (interface{}, error) {
	p := node.Next
	for j := 1; j < i; j++ {
		if p == nil {
			return "", errors.New("没有相应的节点～～")
		}
		p = p.Next
	}

	return p.Data, nil
}

// 插入到第i个元素
func (node *Node) Insert(i int, value interface{}) error {
	p := node
	for j := 1; j < i; j++ {
		if p == nil {
			return errors.New("没有相应的节点～～")
		}
		p = p.Next
	}
	s := new(Node)
	s.Data = value
	s.Next = p.Next
	p.Next = s
	return nil
}

// 删除第i个元素
func (node *Node) Delete(i int) error {
	p := node
	for j := 1; j < i; j++ {
		if p == nil {
			return errors.New("没有相应的节点～～")
		}
		p = p.Next
	}
	q := p.Next
	p.Next = q.Next
	return nil
}

// 创建链表，从头部插入
func (node *Node) InsertWithHead(n int){
	node.Next = nil
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++  {
		p := new(Node)
		//p.Data = rand.Int()
		p.Data = i
		p.Next = node.Next
		node.Next = p
	}
}

// 创建链表，从尾部插入
func (node *Node) InsertWithTail(n int){
	node.Next = nil
	p := node
	//rand.Seed(time.Now().Unix())
	rand.Seed(100)
	for i:=0; i< n ; i++  {
		q := new(Node)
		q.Data = i
		p.Next = q
		p = q
	}
}

// 获取全值
func (node *Node) List()  {
	p := node
	for {
		if p == nil {
			break
		}
		fmt.Println(p)
		p = p.Next
	}
}

// 清除链表
func (node *Node) Clear() {
	p := node
	for {
		if p.Next == nil {
			break
		}
		q := p.Next
		p.Next = nil
		p.Data = 0
		p = q
	}
}

// 测试接口
func (node *Node) testThisNode(contract NodeContract){
	contract.Insert(1)
	contract.Init()
	contract.Length()
	contract.Get(1)
	//node.Delete()
}

