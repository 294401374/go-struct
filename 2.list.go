package main

import (
	"errors"
	"fmt"
)

/*
线性表的基本方法
1、初始化
func Init()

2.判断空
func IsEmpty()

3.清除
func ClearList()

4.获取元素
func GetElement()

5.删除元素
func DeleteElement()

6. 获取长度
func ListLength()

7. 判断第i元素是否存在
func ElementExist()

8. 插入元素
func InsertElement()

// 有些还有其他的操作，例如：

8.合并
func Union()

9.交集
func Intersection()

*/

const MAXSIZE = 5

// 线性表顺序存储结构
type List struct {
	Element [MAXSIZE]int
	Length  int
}

// 获取第i个元素
func (L *List) GetElement(i, value int) (int, error) {
	if i >= L.Length {
		return 0, errors.New("i不可以超过List的长度")
	}
	return L.Element[i], nil
}

// 插入元素
// 线性表插入注意事项：
// 1、表长度不能超过数组长度
// 2、从最后一个元素往前遍历到第i元素并往后移动一个元素
// 3、如果插入的位置不合理抛出一场
// 4、将元素插入到i
// 5、表长度+1
func (L *List) InsertElement(i, value int) error {
	if L.Length >= MAXSIZE {
		return errors.New("List长度不能超过MAXSIZE")
	}

	if i < 0 || i > L.Length {
		return errors.New("i不在正确的范围")
	}
	if i <= L.Length-1 {
		for j := L.Length - 1; j >= i; j-- {
			L.Element[j+1] = L.Element[j]
		}
	}
	L.Element[i] = value
	L.Length++
	return nil
}

func (L *List) DelElement(i int) error {
	if L.Length == 0 {
		return errors.New("L为空")
	}
	if i < 0 || i > L.Length-1 {
		return errors.New("i 超出规定范围")
	}
	L.Element[i] = 0
	if i < L.Length-1 {
		for j := i + 1; j <= L.Length-1; j++ {
			L.Element[j-1] = L.Element[j]
			L.Element[j] = 0
		}
	}
	L.Length--
	return nil

}

func main() {
	list := &List{}
	fmt.Println(list)
	list.InsertElement(0, 12)
	list.InsertElement(1, 1)
	list.InsertElement(2, 3)
	list.InsertElement(1, 13)
	list.InsertElement(4, 14)
	err := list.InsertElement(4, 14)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
	// 删除
	list.DelElement(1)
	fmt.Println(list)
}
