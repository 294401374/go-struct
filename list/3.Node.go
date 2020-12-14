package list

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

// 创建一个链表
func (node *Node) Init(i int) {
	rand.Seed(time.Now().Unix())
	for j := 0; j < i; j++ {
		node.Insert(j, rand.Intn(100))
	}
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
func (node *Node) InsertWithHead(n int) {
	node.Next = nil
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		p := new(Node)
		//p.Data = rand.Int()
		p.Data = i
		p.Next = node.Next
		node.Next = p
	}
}

// 创建链表，从尾部插入
func (node *Node) InsertWithTail(n int) {
	node.Next = nil
	p := node
	//rand.Seed(time.Now().Unix())
	rand.Seed(100)
	for i := 0; i < n; i++ {
		q := new(Node)
		q.Data = i
		p.Next = q
		p = q
	}
}

// 获取全值
func (node *Node) List() {
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

// 翻转链表--递归(未完成)
func ReverseWithRecurse(head *Node) *Node {
	if head.Next == nil || head == nil {
		return head
	} else {
		newNode := new(Node)
		newNode.Next = ReverseWithRecurse(head)
		return newNode
	}
}

// 翻转链表--循环
func (node *Node) Reverse() {
	//  先声明两个变量
	//  前一个节点
	preNode := new(Node)
	//  后一个节点
	nextNode := new(Node)
	for node.Next != nil {
		//  保存头节点的下一个节点，
		nextNode = node.Next
		//  将头节点指向前一个节点
		node.Next = preNode
		//  更新前一个节点
		preNode = node
		//  更新头节点
		node = nextNode
	}
}

// 没懂！！！！！！！！！！！
// 最后一个的指针和data都是nil,所以错误
func (node *Node)ReverseNode(head *Node) *Node {
	//  先声明两个变量
	//  前一个节点
	var preNode *Node
	preNode = nil
	//  后一个节点
	nextNode := new(Node)
	nextNode = nil
	for head != nil {
		//  保存头节点的下一个节点，
		nextNode = head.Next
		//  将头节点指向前一个节点
		head.Next = preNode
		//  更新前一个节点
		preNode = head
		//  更新头节点
		head = nextNode
	}
	headNode := new(Node)
	headNode.Next = preNode
	return headNode
}

// 测试接口
func (node *Node) testThisNode(contract NodeContract) {
	contract.Insert(1)
	contract.Init()
	contract.Length()
	contract.Get(1)
	//node.Delete()
}
