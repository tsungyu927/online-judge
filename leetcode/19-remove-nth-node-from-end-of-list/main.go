/*
Problem: Remove Nth Node From End of List
Link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
*/
package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
  slow, fast := dummy, dummy

  // Move fast to n+1 pos
  for i := 0; i < n+1; i++ {
    fast = fast.Next
  }

  // Move both slow and fast step by step (until fast to the end)
  for fast != nil {
    slow = slow.Next
    fast = fast.Next
  }

  // Let the next of prev of n point to the next of n
  slow.Next = slow.Next.Next

  return dummy.Next
}