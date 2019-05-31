package main

import (
	"Common"
	"LogicFun"
)

func main() {
	lists := make([]*Common.ListNode, 0)
	// list1 := &Common.ListNode{
	// 	Val: 1,
	// }
	// list1.Next = &Common.ListNode{
	// 	Val: 3,
	// }
	// list1.Next.Next = &Common.ListNode{
	// 	Val: 4,
	// }
	// list2 := &Common.ListNode{
	// 	Val: 1,
	// }
	// list2.Next = &Common.ListNode{
	// 	Val: 4,
	// }
	// list2.Next.Next = &Common.ListNode{
	// 	Val: 5,
	// }
	// list3 := &Common.ListNode{
	// 	Val: 2,
	// }
	// list3.Next = &Common.ListNode{
	// 	Val: 6,
	// }

	// lists = append(lists, list1, list2, list3)
	LogicFun.MergeKLists(lists)
}
