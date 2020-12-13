package List

import (
	"errors"
	"fmt"
)

// 1、保存棋盘到文件中
// 2、读取文件并重现

func chess(column, row int) {
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			//fmt.Printf("第%d列\t", j)
			if i == 1 && j == 1 {
				fmt.Printf("%d\t", 1)
			} else if i == 2 && j == 3 {
				fmt.Printf("%d\t", 3)
			} else {
				fmt.Printf("%d\t", 0)
			}
		}
		//fmt.Printf("第%d行   \t\n", i)
		fmt.Println()
	}
}

/**
棋盘记录器
*/
type chesses struct {
	maxSize int
	array   [5][]int
	value   int
}

// 设置稀疏数组值
func (che *chesses) set(column, row, value int) error {
	if row > che.maxSize {
		return errors.New("行数不能大于最大值")
	}
	che.array[row][column] = value
	return nil
}

// 获取稀疏数组值
func (che *chesses) get(column, row, value int) (int, error) {
	if row > che.maxSize {
		return 0, errors.New("行数不能大于最大值")
	}
	return che.array[row][column], nil
}

type chessNode struct {
	row int
	column int
	value int
}

