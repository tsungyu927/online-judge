package main

fun main() {
    val rawData = intArrayOf(1, 2, 6, 3, 4, 5, 6)
    val head = createLinkedList(rawData)
    val solution = Solution()
    val result = solution.removeElements(head, 6)

    printLinkedList(result)
}

fun createLinkedList(nums: IntArray): ListNode? {
    if (nums.isEmpty()) return null

    val head = ListNode(nums[0])
    var current = head

    for (i in 1 until nums.size) {
        current.next = ListNode(nums[i])
        current = current.next!! // Move current to the new node
    }

    return head
}

fun printLinkedList(node: ListNode?) {
    var current = node
    while (current != null) {
        print("${current.`val`} -> ")
        current = current.next
    }
    println("null")
}

class ListNode(
    var `val`: Int,
) {
    var next: ListNode? = null
}

class Solution {
    fun removeElements(
        head: ListNode?,
        `val`: Int,
    ): ListNode? {
        val dummy = ListNode(0).apply { next = head }
        var current = dummy

        while (current.next != null) {
            if (current.next!!.`val` == `val`) {
                current.next = current.next?.next
            } else {
                current = current.next!!
            }
        }

        return dummy.next
    }
}
