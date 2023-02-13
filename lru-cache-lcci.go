package leetcode

import "container/list"

// 问题: LRU (最近最少使用)
//
// 实现:
// 1) hash 判断 key 是否存在, 提供 O(1) 查找能力;
// 2) list 提供 排序能力. 最近使用的元素, 排列在链表头部;
//
// 注意事项:
// 1) get 影响元素使用, 需要调整链表;
// 2) set 需要区分两种情况:
//    a) 元素存在: 1) 覆盖 value; 2) 调整链表结构;
//    b) 元素不存在: 1) 插入链表元素; 2) 维护 hash 查找结构;
//

type (
	LRUCache struct {
		hash     map[int]*list.Element // key type := int
		list     *list.List
		capacity int
	}
	pair struct {
		key   int
		value int
	}
)

func ConstructorLRU(capacity int) LRUCache {
	return LRUCache{
		hash:     make(map[int]*list.Element, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.hash[key]
	if ok {
		this.list.MoveToFront(e)
		return e.Value.(*pair).value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// key 存在
	e, ok := this.hash[key]
	if ok {
		e.Value.(*pair).value = value
		this.list.MoveToFront(e)
		return
	}

	// key 不存在
	e = this.list.PushFront(
		&pair{key: key, value: value},
	)
	this.hash[key] = e

	// 容量检查
	if this.list.Len() > this.capacity {
		e = this.list.Back()
		this.list.Remove(e)
		delete(this.hash, e.Value.(*pair).key)
	}
}
