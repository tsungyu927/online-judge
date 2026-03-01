package main

class Solution {
    fun detectCycle(head: ListNode?): ListNode? {
        var slow: ListNode? = head
        var fast: ListNode? = head

        // Check whether there exist a cycle
        while (true) {
            if (fast == null || fast.next == null) {
                return null
            }

            slow = slow?.next
            fast = fast?.next?.next

            if (slow == fast) break
        }

        // There exist a cycle, find the entrance point
        slow = head
        while (slow != fast) {
            slow = slow?.next
            fast = fast?.next
        }

        return slow
    }
}
