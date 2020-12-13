package List

type NodeContract interface {
	// 初始化
	Init()
	// 插入
	Insert(i int) error
	// 删除
	Delete(i int) error
	// 获取某个节点
	Get(i int) (int, error)
	// 获取链表长度
	Length() int
}