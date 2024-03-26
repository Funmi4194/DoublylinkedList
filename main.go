package main

import (
	doublyLinkkedList "github.com/funmi4194/doublylinkedlist"
)

func main() {
	d := doublyLinkkedList.InitializeDoublyLinkedList()

	var s doublyLinkkedList.DoublyLinkedL = d
	s.InsertFirst("Funmi")
	s.InsertFirst("tayo")
	s.InsertAfter("tayo", "bunmi")
	s.InsertAfter("bunmi", "precious")
	s.InsertAfter("tayo", "ayo")

	s.PrintList()

	s.DeleteInner("ayo")
	s.PrintList()
}
