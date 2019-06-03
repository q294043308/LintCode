package Common

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitList(vals []int) *ListNode {
	var node, list *ListNode
	for _, val := range vals {
		if node == nil {
			node = &ListNode{
				Val: val,
			}
			list = node
		} else {
			node.Next = &ListNode{
				Val: val,
			}
			node = node.Next
		}
	}
	return list
}
