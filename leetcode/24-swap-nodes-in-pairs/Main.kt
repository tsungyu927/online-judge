package main

class Solution {
    fun swapPairs(head: ListNode?): ListNode? {
        if (head?.next == null) {
            return head
        }

        val first = head
        val second = head.next
        val next = second?.next
        second?.next = first
        first?.next = swapPairs(next)

        return second
    }
}
