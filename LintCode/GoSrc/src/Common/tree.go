package Common

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 173. Binary Search Tree Iterator
type BSTIterator struct {
	cur int
	arr []int
}

func TreeConstructor(root *TreeNode) BSTIterator {
	res := BSTIterator{}
	res.Build(root)
	return res
}

func (this *BSTIterator) Build(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		this.Build(root.Left)
	}
	this.arr = append(this.arr, root.Val)
	if root.Right != nil {
		this.Build(root.Right)
	}
}

func (this *BSTIterator) Next() int {
	if this.cur >= len(this.arr) {
		return -1
	}
	res := this.arr[this.cur]
	this.cur++
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.cur < len(this.arr)
}
