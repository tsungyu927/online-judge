package main

class Solution {
    fun removeNthFromEnd(
        head: ListNode?,
        n: Int,
    ): ListNode? {
        var slow = head
        var fast = head

        // Initialize the fast pointer
        for (i in 1..n) {
            fast = fast?.next
        }
        if (fast == null) return head?.next

        while (fast?.next != null) {
            slow = slow?.next
            fast = fast?.next
        }

        slow?.next = slow?.next?.next

        return head
    }
}
