package HuffmanTree

type HuffmanNode struct {
	Element string
	Weight  int
	Left    *HuffmanNode
	Right   *HuffmanNode
}

func NewHuffmanNode(element string, weight int) *HuffmanNode {
	return &HuffmanNode{
		Element: element,
		Weight:  weight,
	}
}

func (node *HuffmanNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func MergeNodes(first, second *HuffmanNode) *HuffmanNode {
	return &HuffmanNode{
		Weight: first.Weight + second.Weight,
		Left:   first,
		Right:  second,
	}
}
