package tree

import (
	"slices"
)

// BPTreeDegree B+树的阶数
const BPTreeDegree = 5

// BPTreeNode b+树节点，通过内存模拟b+树的实现
type BPTreeNode struct {
	IsLeaf   bool          // 是否为叶子节点
	Parent   *BPTreeNode   // 父节点
	Keys     []int         // 中间节点键
	Children []*BPTreeNode // 子节点
	Next     *BPTreeNode   // 叶子节点链表
}

// NewBPTree 新建b+树，此处返回的是跟节点
func NewBPTree() *BPTreeNode {
	return &BPTreeNode{
		IsLeaf:   true,
		Parent:   nil,
		Keys:     make([]int, 0),
		Children: make([]*BPTreeNode, 0),
		Next:     nil,
	}
}

// Insert 插入键值对，如果遇到相同的键，则执行更新操作
func (b *BPTreeNode) Insert(key int, data string) {
	node := b.searchLeaf(key, b)
	if len(node.Keys) <= 0 {
		node.Keys = append(node.Keys, key)
		return
	}
	pos, _ := slices.BinarySearch(node.Keys, key)
	keys := make([]int, 0)
	keys = append(keys, node.Keys[:pos]...)
	keys = append(keys, key)
	keys = append(keys, node.Keys[pos:]...)
	node.Keys = keys

	if len(node.Keys) >= BPTreeDegree {
		b.splitNode(node)
	}
	return
}

func (b *BPTreeNode) newNode() *BPTreeNode {
	return &BPTreeNode{
		IsLeaf:   false,
		Parent:   nil,
		Keys:     make([]int, 0),
		Children: make([]*BPTreeNode, 0),
		Next:     nil,
	}
}

func (b *BPTreeNode) insertParent(left *BPTreeNode, key int, right *BPTreeNode) {
	parent := left.Parent
	if parent == nil { // 没有parent，向上增长
		node := b.newNode()
		node.IsLeaf = left.IsLeaf
		node.Keys = left.Keys
		node.Children = left.Children
		node.Next = left.Next

		left.Parent = nil
		left.Keys = make([]int, 0)
		left.Children = make([]*BPTreeNode, 0)
		left.Next = nil
		left.IsLeaf = false
		left.Keys = append(left.Keys, key)
		left.Children = append(left.Children, node, right)

		node.Parent = left
		right.Parent = left
		return
	}

	pos, _ := slices.BinarySearch(parent.Keys, key)
	keys := make([]int, 0)
	keys = append(keys, parent.Keys[:pos]...)
	keys = append(keys, key)
	keys = append(keys, parent.Keys[pos:]...)
	parent.Keys = keys

	children := make([]*BPTreeNode, 0)
	pos = 0
	for i := 0; i < len(parent.Children); i++ {
		pos = i
		children = append(children, parent.Children[i])
		if left == parent.Children[i] {
			break
		}
	}
	children = append(children, right)
	for i := pos + 1; i < len(parent.Children); i++ {
		children = append(children, parent.Children[i])
	}
	parent.Children = children

	if len(parent.Keys) >= BPTreeDegree {
		b.splitNode(parent)
	}
}

func (b *BPTreeNode) splitLeaf(curNode *BPTreeNode) {
	if !curNode.IsLeaf {
		return
	}

	pos := len(curNode.Keys) / 2
	rightNode := b.newNode()
	rightNode.IsLeaf = true
	rightNode.Keys = curNode.Keys[pos:]
	rightNode.Parent = curNode.Parent
	curNode.Keys = curNode.Keys[:pos]
	curNode.Next = rightNode
	b.insertParent(curNode, rightNode.Keys[0], rightNode)
	return
}

func (b *BPTreeNode) splitInternal(curNode *BPTreeNode) {
	if curNode.IsLeaf {
		return
	}
	pos := len(curNode.Keys) / 2
	rightNode := b.newNode()
	rightNode.IsLeaf = false
	rightNode.Keys = curNode.Keys[pos+1:]
	rightNode.Children = curNode.Children[pos+1:]
	rightNode.Parent = curNode.Parent
	key := curNode.Keys[pos]
	curNode.Keys = curNode.Keys[:pos]
	curNode.Children = curNode.Children[:pos+1]
	b.insertParent(curNode, key, rightNode)
	return
}

func (b *BPTreeNode) splitNode(curNode *BPTreeNode) {
	if len(curNode.Keys) < BPTreeDegree {
		return
	}
	if curNode.IsLeaf {
		b.splitLeaf(curNode)
		return
	}
	b.splitInternal(curNode)
	return
}

func (b *BPTreeNode) Delete(key int) {

}

func (b *BPTreeNode) mergeNode() {

}

func (b *BPTreeNode) searchLeaf(key int, node *BPTreeNode) *BPTreeNode {
	if node.IsLeaf {
		return node
	}
	for i := 0; i < len(node.Keys); i++ {
		if key < node.Keys[i] {
			return b.searchLeaf(key, node.Children[i])
		}
	}
	return b.searchLeaf(key, node.Children[len(node.Children)-1])
}

// Print 打印出整个树的结构，用于调试
func (b *BPTreeNode) Print() {

}
