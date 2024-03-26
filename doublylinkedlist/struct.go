package doublyLinkkedList

type Node struct {
	// Prev holds the address of the previous node
	Prev *Node

	// Data stores the data item
	Data string

	// Next holds the address of the next node
	Next *Node
}

//
type DoublyLinkedList struct {
	// First is the connection link to the first link
	First *Node

	// Last is the connection link to the last link
	Last *Node
}
