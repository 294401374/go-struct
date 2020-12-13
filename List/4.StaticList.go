package List

import (
	"errors"
)

const StaticMaxSize = 10

// 静态链表
// 静态链表是没有指针的，所以用数组形式来实现链表。
// 静态链表的特点:
//		数组每个节点都会有一个cursor（游标），充当指针的作用
//		是用头和尾部节点的cursor分别指向剩余的可用节点和元素的头节点
type StaticNode struct {
	pos    int
	Data   string
	Cursor int
}

type StaticList struct {
	List [StaticMaxSize]StaticNode
}

// 初始化
// cursor 相当于指针
// 设置第一个元素的游标为下一个可用空余数据（指针）
// 设置最后一个元素的游标为指向第一个元素的游标（指针）
func (StList *StaticList) InitList(n int) error {
	if n < 3 {
		return errors.New("元素不能小于3")
	}
	// 数组
	StList.List = [StaticMaxSize]StaticNode{}
	for i := 0; i < StaticMaxSize-1; i++ {
		// 设置cursor
		StList.List[i].Cursor = i + 1
		StList.List[i].pos = i + 1
	}
	//设置最后一个元素的cursor
	StList.List[StaticMaxSize-1].Cursor = 0
	return nil
}

func (StList *StaticList) CreateNode() int {
	i := StList.List[0].Cursor
	if i != 0 {
		StList.List[0].Cursor = StList.List[i].Cursor
	}
	return i
}

// 添加元素
// 特点：
//	1、创建元素，第一个元素的cursor改成新元素的cursor
// 	2、创建一个元素，获取上个元素的cursor并赋给新增的元素把，把上个元素的cursor改成自己
func (StList *StaticList) InsertElement(i int, value string) error {
	var n, p int
	p = StaticMaxSize - 1
	if i < 0 || i > StList.ListLength()+1 {
		return errors.New("i不符合规范")
	}
	n = StList.CreateNode()
	if n != 0 {
		StList.List[n].Data = value
		// 从头节点开始遍历
		for j := 1; j <= i-1; j++ {
			// 从第一个节点开始遍历
			// 获取第i-1个的Cursor
			p = StList.List[p].Cursor
		}
		StList.List[n].Cursor = StList.List[p].Cursor
		StList.List[p].Cursor = n
		return nil
	}
	return errors.New("发生错误")
}

// 获取链表长度
func (StList *StaticList) ListLength() int {
	p := StaticMaxSize - 1
	var i int
	for {
		if StList.List[p].Cursor == 0 {
			break
		}
		p = StList.List[p].Cursor
		i++
	}
	return i
}

// 释放链表节点
func (StList *StaticList) FreeElement(i int) {
	StList.List[i].Cursor = StList.List[0].Cursor
	StList.List[0].Cursor = i
}

// 删除节点
//
func (StList *StaticList) DeleteNode(i int) error {
	if i < 0 || i > StList.ListLength()+1 {
		return errors.New("i 不规范")
	}
	p := StaticMaxSize - 1
	for j := 0; j < i-1; j ++ {
		p = StList.List[p].Cursor
	}
	StList.List[p].Cursor = StList.List[i].Cursor
	StList.FreeElement(i)
	return nil
}

