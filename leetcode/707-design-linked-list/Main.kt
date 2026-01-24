package main

class MyLinkedList {
    private var head: Node? = null
    private var length = 0

    fun get(index: Int): Int = findByIndex(index)?.value ?: -1

    fun addAtHead(`val`: Int) {
        addAtIndex(0, `val`)
    }

    fun addAtTail(`val`: Int) {
        addAtIndex(length, `val`)
    }

    fun addAtIndex(
        index: Int,
        `val`: Int,
    ) {
        if (index > length) return

        val prevNode = findByIndex(index - 1)
        val node = Node(`val`)

        when {
            prevNode == null -> {
                head = node
            }

            index == 0 -> {
                node.next = prevNode
                head = node
            }

            else -> {
                node.next = prevNode.next
                prevNode.next = node
            }
        }
        length++
    }

    fun deleteAtIndex(index: Int) {
        if (length == 0 || index > length - 1 || index < 0) return
        if (index == 0) {
            head = head?.next
            length--
            return
        }

        val prevNode = findByIndex(index - 1)
        prevNode?.next = prevNode?.next?.next
        length--
    }

    private fun findByIndex(index: Int): Node? {
        if (length == 0 || index > length - 1) return null
        if (index <= 0) return head

        var next: Node? = head
        for (i in 0 until index) {
            next = next?.next
        }
        return next
    }

    private data class Node(
        val value: Int,
        var next: Node? = null,
    )
}
