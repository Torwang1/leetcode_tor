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
	if e, ok := this.search[key]; ok {
		this.list.MoveToFront(e)
		return e.Value.(pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	item := pair{key: key, value: value}

	if e, ok := this.search[key]; ok {
		e.Value = item
		this.list.MoveToFront(e)
		return
	}

	e := this.list.PushFront(item)
	this.search[key] = e

	if this.list.Len() > this.capacity {
		if e := this.list.Back(); e != nil {
			item := this.list.Remove(e)
			delete(this.search, item.(pair).key)
		}
	}
}
