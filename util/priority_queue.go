package util

import (
	"container/heap"
)

// Item 是优先队列中包含的元素。
type Item struct {
	// 元素的值
	value    interface{}
	// 元素在队列中的优先级。
	priority int
	// 元素在堆中的索引。可以用于更新操作，它由 heap.Interface 定义的方法维护。
	index int
}

// PriorityQueue 优先队列
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 希望Pop 返回的是最大值，因此这里使用大于号进行对比。
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// update 更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}