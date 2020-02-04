package main

import (
	"Common"
)

func main() {
	cache := Common.Constructor(3 /* capacity */)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)       // evicts key 2
	cache.Put(4, 4)       // evicts key 1
	println(cache.Get(4)) // returns 1
	println(cache.Get(3)) // returns -1 (not found)
	println(cache.Get(2)) // returns -1 (not found)
	println(cache.Get(1)) // returns 3
	cache.Put(5, 5)       // evicts key 1
	println(cache.Get(1)) // returns 4
	println(cache.Get(2)) // returns 4
	println(cache.Get(3)) // returns 4
	println(cache.Get(4)) // returns 4
	println(cache.Get(5)) // returns 4
	// head := &Common.ListNode{Val: 1}
	// head.Next = &Common.ListNode{Val: 2}
	// head.Next.Next = &Common.ListNode{Val: 3}
	// head.Next.Next.Next = &Common.ListNode{Val: 4}
	// head.Next.Next.Next.Next = &Common.ListNode{Val: 5}
	// // head.Left = &Common.TreeNode{Val: 2}
	// // head.Right = &Common.TreeNode{Val: 3}
	// // head.Left.Left = &Common.TreeNode{Val: 3}
	// // head.Left.Right = &Common.TreeNode{Val: 4}
	// // head.Right.Left = &Common.TreeNode{Val: 15}
	// // head.Right.Right = &Common.TreeNode{Val: 7}
	// // head.Left.Right.Left = &Common.TreeNode{Val: 31}
	// LogicFun.ReorderList(head)
}
