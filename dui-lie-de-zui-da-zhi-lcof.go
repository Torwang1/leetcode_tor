package leetcode

import "container/list"

type MaxQueue struct {
	elements *list.List
	maxs     *list.List
}

func ConstructorII() MaxQueue {
	return MaxQueue{
		elements: list.New(),
		maxs:     list.New(),
	}
}

func (this *MaxQueue) Max_value() int {
	e := this.maxs.Front()
	if e == nil {
		return -1
	}
	return e.Value.(int)
}

// Push_back 把 maxs 队列中, 小于 value 的元素弹出.
func (this *MaxQueue) Push_back(value int) {
	e := this.maxs.Back()
	for e != nil && e.Value.(int) < value {
		this.maxs.Remove(e)
		e = this.maxs.Back()
	}

	this.elements.PushBack(value)
	this.maxs.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	e := this.elements.Front()
	if e == nil {
		return -1
	}
	this.elements.Remove(e)

	m := this.maxs.Front()
	if m.Value.(int) == e.Value.(int) {
		this.maxs.Remove(m)
	}

	return e.Value.(int)
}
