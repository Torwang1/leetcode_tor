package leetcode

import "container/list"

type (
	LRUCache struct {
		search   map[int]*list.Element
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
		search:   make(map[int]*list.Element, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.search[key]
	if ok {
		this.list.MoveToFront(e)
		return e.Value.(*pair).value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// key 存在
	e, ok := this.search[key]
	if ok {
		e.Value.(*pair).value = value
		this.list.MoveToFront(e)
		return
	}

	// key 不存在
	e = this.list.PushFront(
		&pair{key: key, value: value},
	)
	this.search[key] = e

	// 容量检查
	if this.list.Len() > this.capacity {
		e = this.list.Back()
		this.list.Remove(e)
		delete(this.search, e.Value.(*pair).key)
	}
}
