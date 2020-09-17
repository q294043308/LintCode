package main

import logicfun "Logicfun"

func main() {
	// head := &Common.ListNode{Val: 4}
	// head.Next = &Common.ListNode{Val: 1}
	// head.Next.Next = &Common.ListNode{Val: 8}
	// head.Next.Next.Next = &Common.ListNode{Val: 4}

	// head2 := &Common.ListNode{Val: 2}
	// head2.Next = head.Next
	// head.Next.Next.Next.Next = &Common.ListNode{Val: 0}
	// // head.Left = &Common.TreeNode{Val: 2}
	// // head.Right = &Common.TreeNode{Val: 3}
	// // head.Left.Left = &Common.TreeNode{Val: 3}
	// // head.Left.Right = &Common.TreeNode{Val: 4}
	// // head.Right.Left = &Common.TreeNode{Val: 15}
	// // head.Right.Right = &Common.TreeNode{Val: 7}
	// close(test)

	print(logicfun.TrailingZeroes(25))
}
