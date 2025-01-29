package main

import (
	"testing"
)

// Test MyLinkedList implementation
func TestMyLinkedList(t *testing.T) {
	list := Constructor()

	// Test AddAtHead
	list.AddAtHead(1)
	if list.Get(0) != 1 {
		t.Errorf("Expected 1 at index 0, got %d", list.Get(0))
	}

	// Test AddAtTail
	list.AddAtTail(3)
	if list.Get(1) != 3 {
		t.Errorf("Expected 3 at index 1, got %d", list.Get(1))
	}

	// Test AddAtIndex
	list.AddAtIndex(1, 2) // List should be [1,2,3]
	if list.Get(1) != 2 {
		t.Errorf("Expected 2 at index 1, got %d", list.Get(1))
	}

	// Test Get method
	if list.Get(0) != 1 {
		t.Errorf("Expected 1 at index 0, got %d", list.Get(0))
	}
	if list.Get(1) != 2 {
		t.Errorf("Expected 2 at index 1, got %d", list.Get(1))
	}
	if list.Get(2) != 3 {
		t.Errorf("Expected 3 at index 2, got %d", list.Get(2))
	}
	if list.Get(3) != -1 {
		t.Errorf("Expected -1 at index 3 (out of bounds), got %d", list.Get(3))
	}

	// Test DeleteAtIndex (deleting middle element)
	list.DeleteAtIndex(1) // List should become [1,3]
	if list.Get(1) != 3 {
		t.Errorf("Expected 3 at index 1 after deletion, got %d", list.Get(1))
	}

	// Test DeleteAtIndex (deleting head)
	list.DeleteAtIndex(0) // List should become [3]
	if list.Get(0) != 3 {
		t.Errorf("Expected 3 at index 0 after deleting head, got %d", list.Get(0))
	}

	// Test DeleteAtIndex (deleting last element)
	list.DeleteAtIndex(0) // List should become empty
	if list.Get(0) != -1 {
		t.Errorf("Expected -1 at index 0 (empty list), got %d", list.Get(0))
	}

	// Test DeleteAtIndex on empty list
	list.DeleteAtIndex(0) // Should not crash or change anything
	if list.Get(0) != -1 {
		t.Errorf("Expected -1 at index 0 on empty list, got %d", list.Get(0))
	}

	// Test AddAtIndex with index out of bounds
	list.AddAtIndex(5, 10) // Should not insert
	if list.Get(0) != -1 {
		t.Errorf("Expected -1 since index 5 is out of bounds, got %d", list.Get(0))
	}

	// Edge case: Add at index 0 when list is empty
	list.AddAtIndex(0, 42) // Should insert 42 at head
	if list.Get(0) != 42 {
		t.Errorf("Expected 42 at index 0, got %d", list.Get(0))
	}
}
