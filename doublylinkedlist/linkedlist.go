package doublyLinkkedList

import (
	"fmt"
)

type DoublyLinkedL interface {
	PrintList()
	InsertFirst(data string)
	InsertEnd(data string)
	InsertAfter(prevInsertData, data string)
	DeleteFirst()
	DeleteLast()
	DeleteInner(data string)
	DeleteWithParameter(data string)
}

func InitializeDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// transverse from the head of the list
// Function to print the doubly linked list
func (d *DoublyLinkedList) PrintList() {
	// current points to the first nide of the doubly link list
	current := d.First
	//iterates through the link list if its not equals nil
	for current != nil {
		fmt.Printf("(%s) ", current.Data)
		// update the pointer to point to the next node in the list
		current = current.Next
	}
}

// InsertFirst adds a node with three compartment at the beginning of doublylinkedlist
func (d *DoublyLinkedList) InsertFirst(data string) {
	newNode := &Node{Data: data}

	// if no node exist before return newNode as first and last node
	if d.First == nil {
		d.First = newNode
		d.Last = newNode
	} else {

		// point the next of newNode to the old first node
		newNode.Next = d.First
		// point the prev of new node to nil
		newNode.Prev = nil
		// update first prev node
		// point prev of the first node to the new node (this makes first node become second)
		d.First.Prev = newNode
		//point first to the new node
		d.First = newNode

	}
}

func (d *DoublyLinkedList) InsertAfter(prevInsertData, data string) {

	var prevNode *Node
	current := d.First

	// get the previous data
	for current != nil {
		if current.Data == prevInsertData {
			prevNode = current
			break
		}
		current = current.Next
	}

	newNode := &Node{Data: data}

	// set next of newnode to next from previous node
	newNode.Next = prevNode.Next
	// set next of previous node to address of newnode
	prevNode.Next = newNode

	// set the prev of new node to prev node
	newNode.Prev = prevNode

	// set the previous of next node (newnode's next) to newnode
	if newNode.Next != nil {
		newNode.Next.Prev = newNode
	}
}

func (d *DoublyLinkedList) InsertEnd(data string) {
	newNode := &Node{Data: data}

	newNode.Next = nil

	//assign nil to the next of newnode
	newNode.Next = nil

	if d.First == nil {
		d.First = newNode
		d.Last = newNode
		newNode.Prev = nil
		return
	}

	current := d.First

	// if the link is not empty, loop through to the end of the list
	for current.Next != nil {
		current = current.Next
	}

	// set nwnode to the next of current node
	current.Next = newNode

	// set the last node to the previous of newnode
	newNode.Prev = current
}

func (d *DoublyLinkedList) DeleteFirst() {
	deleteNode := &Node{Prev: nil, Next: nil}

	// if del_node is the first node, point the head pointer to the next of deleteNode
	deleteNode = d.First

	d.First = deleteNode.Next
	deleteNode.Next.Prev = nil
}

func (d *DoublyLinkedList) DeleteLast() {
	deleteNode := &Node{Prev: nil, Next: nil}

	// if del_node is the last node, point the last pointer to the prev of deleteNode
	deleteNode = d.Last

	d.Last = deleteNode.Prev
	deleteNode.Prev.Next = nil
}

func (d *DoublyLinkedList) DeleteInner(data string) {
	deleteNode := &Node{Data: data}

	current := d.First

	for current != nil {
		if current.Data == data {
			deleteNode = current
			fmt.Println(current.Prev.Data)
			deleteNode.Prev.Next = deleteNode.Next
			deleteNode.Next.Prev = deleteNode.Prev
			break
		}
		current = current.Next
	}
}

func (d *DoublyLinkedList) DeleteWithParameter(data string) {

	current := d.First
	for current != nil {
		if current.Data == data {
			if current == d.First {
				d.First = current.Next
				if d.First != nil {
					d.First.Prev = nil
				}
			} else if current == d.Last {
				d.Last = current.Prev
				if d.Last != nil {
					d.Last.Next = nil
				}
			} else {
				current.Prev.Next = current.Next
				current.Next.Prev = current.Prev
			}
			return
		}
		current = current.Next
	}
}
