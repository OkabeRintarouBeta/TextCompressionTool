package HuffmanTree

import (
	"container/heap"
)

type HuffmanTree struct {
	Root HuffmanNode
}

func BuildHuffmanEncoding(frequencyMap map[string]int) map[string]string {
	nodes := make([]*HuffmanNode, 0)
	for k, v := range frequencyMap {
		nodes = append(nodes, NewHuffmanNode(k, v))
	}
	pq := buildPriorityQueue(nodes)
	tree := buildHuffmanTree(pq)
	encoding := make(map[string]string)
	buildHuffmanEncodingHelper(encoding, &tree.Root, "")
	return encoding
}

func buildHuffmanEncodingHelper(charMap map[string]string, currNode *HuffmanNode, currRep string) {
	if currNode == nil {
		return
	}
	if currNode.IsLeaf() {
		charMap[currNode.Element] = currRep
	}
	buildHuffmanEncodingHelper(charMap, currNode.Left, currRep+"0")
	buildHuffmanEncodingHelper(charMap, currNode.Right, currRep+"1")
}
func buildHuffmanTree(pq *PriorityQueue) *HuffmanTree {
	for pq.Len() > 1 {
		first := heap.Pop(pq).(*Item)
		second := heap.Pop(pq).(*Item)
		newNode := MergeNodes(first.Value, second.Value)
		heap.Push(pq, newNode)
	}
	rootItem := heap.Pop(pq).(*Item)
	return &HuffmanTree{Root: *rootItem.Value}
}
