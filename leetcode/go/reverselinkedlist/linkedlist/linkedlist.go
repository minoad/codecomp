package linkedlist

type LinkedList struct {
	Head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func (n *LinkedList) Add(val int) {
	// create a node with the supplied value.
	node := Node{Val: val}
	// if the head is nil, then we have no entries yet.
	// set it to the head.
	if n.Head == nil {
		n.Head = &node
	} else {
		// if the head is not nil, then we have an entry allready.
		// set the next val on the current head to this node.
		node.Next = n.Head
		// replace the head node with this node.
		n.Head = &node
	}

}
