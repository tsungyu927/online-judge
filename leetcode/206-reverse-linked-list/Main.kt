package main

class Solution {
    fun reverseList(head: ListNode?): ListNode? {
        var currentNode: ListNode? = head
        var prevNode: ListNode? = null

        while (currentNode != null) {
            val nextNode = currentNode.next
            currentNode.next = prevNode
            prevNode = currentNode
            currentNode = nextNode
        }

        return prevNode
    }
}
