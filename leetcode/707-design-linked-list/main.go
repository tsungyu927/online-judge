/*
Problem: Design Linked List
Link: https://leetcode.com/problems/design-linked-list/description/
*/
package main
type Node struct {
  val int
  next *Node
}

type MyLinkedList struct {
  head *Node
  length int
}


func Constructor() MyLinkedList {
  return MyLinkedList{head: nil, length: 0}
}


func (list *MyLinkedList) Get(index int) int {
  currentNode := list.head

  if index > list.length - 1 {
    return -1
  }

  for i := 0; i < index; i++ {
    currentNode = currentNode.next
  }

  return currentNode.val
}


func (list *MyLinkedList) AddAtHead(val int)  {
	head := list.head
  node := Node{val: val, next: head}

	list.head = &node
	list.length++
}


func (list *MyLinkedList) AddAtTail(val int)  {
	node := Node{val: val, next: nil}
	currentNode := list.head

  if list.length == 0 {
    list.head = &node
    list.length++
    return
  }

	for i := 0; i < list.length - 1; i++ {
		currentNode = currentNode.next
	}

	currentNode.next = &node
	list.length++
}


func (list *MyLinkedList) AddAtIndex(index int, val int)  {
	currentNode := list.head

	if index > list.length || index < 0 {
		return
	}

	if index == 0 {
		list.AddAtHead(val)
	} else if index == list.length {
		list.AddAtTail(val)
	} else {
		for i := 0; i < index - 1; i++ {
			currentNode = currentNode.next
		}
		node := Node{val: val, next: currentNode.next}
		currentNode.next = &node
		list.length++
  }
}


func (list *MyLinkedList) DeleteAtIndex(index int)  {
	if index > list.length - 1 || index < 0 {
		return
	}

  if list.length == 0 {
    return
  }

  if list.length == 1 {
    list.head = nil
    list.length--
    return
  }

  if index == 0 {
    list.head = list.head.next
    list.length--
    return
  }

  currentNode := list.head

	for i := 0; i < index - 1; i++ {
		currentNode = currentNode.next
	}

  if index == list.length - 1 {
    currentNode.next = nil
  } else {
	currentNode.next = currentNode.next.next
  }
	list.length--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */