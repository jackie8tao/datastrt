package dst

import "gohelper/errs"

// BPTreeDegree B+树的阶数
const BPTreeDegree = 5

// BPTreeNode b+树节点，通过内存模拟b+树的实现
type BPTreeNode struct {
	IsLeaf   bool          // 是否为叶子节点
	Parent   *BPTreeNode   // 父节点
	Keys     []int         // 中间节点键
	Children []*BPTreeNode // 子节点
	Datas    []string      // 叶子节点数据
	Next     *BPTreeNode   // 叶子节点链表
}

// NewBPTree 新建b+树，此处返回的是跟节点
func NewBPTree() *BPTreeNode {
	return &BPTreeNode{
		IsLeaf:   true,
		Parent:   nil,
		Keys:     make([]int, 0),
		Children: make([]*BPTreeNode, 0),
		Datas:    make([]string, 0),
		Next:     nil,
	}
}

// Insert 插入键值对，如果遇到相同的键，则执行更新操作
func (b *BPTreeNode) Insert(key int, data string) {
	node := b.searchLeaf(key, b)
	if len(node.Keys) <= 0 {
		node.Keys = append(node.Keys, key)
		node.Datas = append(node.Datas, data)
		return
	}

	pos := 0
	for i := 0; i < len(node.Keys); i++ {
		pos = i
		if key < node.Keys[i] {
			break
		}
	}
	pos += 1
	leftKeys, leftDatas := node.Keys[:pos], node.Datas[:pos]
	rightKeys, rightDatas := node.Keys[pos:], node.Datas[pos:]
	node.Keys = append(node.Keys, leftKeys...)
	node.Keys = append(node.Keys, key)
	node.Keys = append(node.Keys, rightKeys...)
	node.Datas = append(node.Datas, leftDatas...)
	node.Datas = append(node.Datas, data)
	node.Datas = append(node.Datas, rightDatas...)
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
		Datas:    make([]string, 0),
		Next:     nil,
	}
}

func (b *BPTreeNode) insertParent(left *BPTreeNode, key int, right *BPTreeNode) {
	parent := left.Parent
	if parent == nil {
		node := b.newNode()
		node.IsLeaf = left.IsLeaf
		node.Keys = left.Keys
		node.Datas = left.Datas
		node.Children = left.Children
		node.Next = left.Next

		left.Parent = nil
		left.Keys = make([]int, 0)
		left.Children = make([]*BPTreeNode, 0)
		left.Datas = make([]string, 0)
		left.Next = nil
		left.IsLeaf = false
		left.Keys = append(left.Keys, key)
		left.Children = append(left.Children, node, right)

		node.Parent = left
		right.Parent = left
		return
	}

	pos := 0
	found := false
	for i := 0; i < len(parent.Keys); i++ {
		if parent.Keys[i] > key {
			pos = i
			found = true
			break
		}
	}
	if !found {
		pos = len(parent.Keys) - 1
	}
	leftKeys := parent.Keys[:pos]
	leftChildren := parent.Children[:pos]
	rightKeys := parent.Keys[pos:]
	rightChildren := parent.Children[pos:]
	parent.Keys = append(leftKeys, key)
	parent.Keys = append(parent.Keys, rightKeys...)
	parent.Children = append(leftChildren, right)
	parent.Children = append(parent.Children, rightChildren...)
	if len(parent.Keys) > BPTreeDegree {
		b.splitInternal(parent)
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
	rightNode.Datas = curNode.Datas[pos:]
	rightNode.Parent = curNode.Parent
	curNode.Keys = curNode.Keys[:pos]
	curNode.Datas = curNode.Datas[:pos]
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

func (b *BPTreeNode) Search(key int) (string, error) {
	node := b.searchLeaf(key, b)
	for i := 0; i < len(node.Keys); i++ {
		if key == node.Keys[i] {
			return node.Datas[i], nil
		}
	}
	return "", errs.ErrNotFound
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

func (b *BPTreeNode) Update(key int, data string) error {
	node := b.searchLeaf(key, b)
	for i := 0; i < len(node.Keys); i++ {
		if key == node.Keys[i] {
			node.Datas[i] = data
			return nil
		}
	}
	return errs.ErrNotFound
}

// Print 打印出整个树的结构，用于调试
func (b *BPTreeNode) Print() {

}
