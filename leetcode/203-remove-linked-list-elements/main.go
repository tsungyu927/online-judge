/*
Problem: Remove Linked List Elements
Link: https://leetcode.com/problems/remove-linked-list-elements/description/
*/
package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	currentNode := dummy

	for currentNode.Next != nil {
		if currentNode.Next.Val == val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}

	return dummy.Next
}