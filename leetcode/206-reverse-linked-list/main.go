/*
Problem: Reverse Linked List
Link:	https://leetcode.com/problems/reverse-linked-list/description/
*/
package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
  var prev *ListNode

  for head != nil {
    next := head.Next
    head.Next = prev
    prev = head
    head = next
  }

  return prev
}