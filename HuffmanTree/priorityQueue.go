package HuffmanTree

import "container/heap"

// Contains the implementation for a min-heap priority queue for huffman nodes

type Item struct {
	Value    *HuffmanNode // The value of the item;
	priority int          // The priority of the item in the queue.
	index    int          // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want pop() to give the lowest
	return pq[i].priority < pq[j].priority
}

func (pq *PriorityQueue) Push(obj any) {

	nodeItem := obj.(*HuffmanNode)
	item := &Item{
		Value:    nodeItem,
		priority: nodeItem.Weight,
		index:    len(*pq),
	}
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	oldQueue := *pq
	originalLen := len(*pq)
	target := oldQueue[originalLen-1]
	*pq = oldQueue[:originalLen-1]
	return target
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) update(item *Item, value *HuffmanNode, priority int) {
	item.Value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func buildPriorityQueue(nodes []*HuffmanNode) *PriorityQueue {
	pq := &PriorityQueue{}
	for _, node := range nodes {
		heap.Push(pq, node) // let Push wrap in Item
	}
	heap.Init(pq)
	return pq
}
