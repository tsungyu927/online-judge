/*
Problem: Swap Nodes in Pairs
Link:	https://leetcode.com/problems/swap-nodes-in-pairs/description/
*/
package main


type ListNode struct {
  Val int
  Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
    return head
  }

  result := head.Next
  head.Next = swapPairs(head.Next.Next)
  result.Next = head

  return result
}