package tree

type BinarySearchTree struct {
	LastResult []int // last result. you need call Clear() before call Order.
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Clear() {
	bst.LastResult = make([]int, 0)
}

func (bst *BinarySearchTree) PreOrder(root *Node) {
	if root == nil {
		return
	}

	return
}

func (bst *BinarySearchTree) InOrder(root *Node) {
	if root == nil {
		return
	}
	return
}

func (bst *BinarySearchTree) PostOrder(root *Node) {
	if root == nil {
		return
	}
	return
}

func (bst *BinarySearchTree) PreOrderByRecursion(root *Node) {
	if root == nil {
		return
	}
	bst.LastResult = append(bst.LastResult, root.Val)
	bst.PreOrderByRecursion(root.Left)
	bst.PreOrderByRecursion(root.Right)
	return
}

func (bst *BinarySearchTree) InOrderByRecursion(root *Node) {
	if root == nil {
		return
	}
	bst.InOrderByRecursion(root.Left)
	bst.LastResult = append(bst.LastResult, root.Val)
	bst.InOrderByRecursion(root.Right)
	return
}

func (bst *BinarySearchTree) PostOrderByRecursion(root *Node) {
	if root == nil {
		return
	}
	bst.PostOrderByRecursion(root.Left)
	bst.PostOrderByRecursion(root.Right)
	bst.LastResult = append(bst.LastResult, root.Val)
	return
}
