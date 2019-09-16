package linkedlist

import (
	"testing"
)

func TestLinkedList_AddToHead(t *testing.T) {
	n := LinkedList{}
	n.Add(5)
	n.Add(6)
	n.Add(7)
	n.Add(8)
}
