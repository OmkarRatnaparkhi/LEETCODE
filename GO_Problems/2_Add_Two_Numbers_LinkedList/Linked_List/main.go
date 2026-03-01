package main

import "fmt"

// NODE represents the structure of a single node - The Structure (The Building Blocks)
// A Node is a box with two compartments:
// Data: The actual number you want to store.
// Next: A pointer (arrow) that points to the next box. If there is no next box, it points to nil (nothing).
type NODE struct {
	Data int
	Next *NODE
}

// SinglyLinkedList represents the linked list class/struct
type SinglyLinkedList struct {
	Head   *NODE
	length int
}

// NewSinglyLinkedList initializes the list (Constructor equivalent)
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{Head: nil}
}

/*
InsertFirst adds a node at the beginning: O(1)
This method puts a new node at the very beginning of the list.
Case A: The list is empty (s.Head == nil)
	1) Create a new node (e.g., 51).
 	2) The Station Master (Head) points directly to this new node.
Case B: The list already has elements
	1) Create a new node (e.g., 21).
	2) Step 1: Point the new node's Next to whatever Head is currently pointing at.
	3) Step 2: Update Head to point to the new node.
This ensures the new node becomes the "First" and the old first node becomes "Second".
*/
func (s *SinglyLinkedList) InsertFirst(no int) {
	newn := &NODE{Data: no, Next: nil}

	if s.Head == nil {
		s.Head = newn
	} else {
		newn.Next = s.Head
		s.Head = newn
	}
}

/*
InsertLast adds a node at the end: O(n)
Case A: The list is empty
	Works exactly like InsertFirst. Head points to the new node.
Case B: The list already has elements
	1) Create the new node.
	2) We start a temporary pointer (temp) at the Head.
	3) The Loop: We move temp forward (temp = temp.Next) until we reach the last node (the one where Next is nil).
	4) The Link: We change that last node’s Next from nil to our new node.
*/
func (s *SinglyLinkedList) InsertLast(no int) {
	newn := &NODE{Data: no, Next: nil}

	if s.Head == nil {
		s.Head = newn
	} else {
		temp := s.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newn
	}
}

// Display prints the elements of the list
func (s *SinglyLinkedList) Display() {
	temp := s.Head
	fmt.Println("Elements of linked list are:")
	for temp != nil {
		fmt.Printf("%d\t", temp.Data)
		temp = temp.Next
	}
	fmt.Println()
}

// Count returns the total number of nodes
func (s *SinglyLinkedList) Count() int {
	iCnt := 0
	temp := s.Head
	for temp != nil {
		iCnt++
		temp = temp.Next
	}
	return iCnt
}

func main() {
	// Object creation
	obj := NewSinglyLinkedList()

	// Method calls
	obj.InsertFirst(51)
	obj.InsertFirst(21)
	obj.InsertFirst(11)
	obj.InsertLast(101)

	obj.Display()

	fmt.Printf("Number of elements: %d\n", obj.Count())
}

/*
package main

import "fmt"

// Node represents a single node in the linked list.
// [T any] makes this a generic struct.
type Node[T any] struct {
	Data T
	Next *Node[T]
}

// SinglyLinkedList represents the actual list structure.
type SinglyLinkedList[T any] struct {
	Head *Node[T]
}

// NewSinglyLinkedList acts like the constructor.
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{Head: nil}
}

// InsertFirst adds a node at the beginning of the list.
func (s *SinglyLinkedList[T]) InsertFirst(no T) {
	newn := &Node[T]{Data: no, Next: nil}

	if s.Head == nil {
		s.Head = newn
	} else {
		newn.Next = s.Head
		s.Head = newn
	}
}

// InsertLast adds a node at the end of the list.
func (s *SinglyLinkedList[T]) InsertLast(no T) {
	newn := &Node[T]{Data: no, Next: nil}

	if s.Head == nil {
		s.Head = newn
	} else {
		temp := s.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newn
	}
}

// Display prints all elements of the linked list.
func (s *SinglyLinkedList[T]) Display() {
	temp := s.Head
	fmt.Println("Elements of linked list are:")
	for temp != nil {
		fmt.Printf("%v\t", temp.Data)
		temp = temp.Next
	}
	fmt.Println()
}

// Count returns the number of nodes in the list.
func (s *SinglyLinkedList[T]) Count() int {
	iCnt := 0
	temp := s.Head
	for temp != nil {
		temp = temp.Next
		iCnt++
	}
	return iCnt
}

func main() {
	// Creating a list of integers
	obj := NewSinglyLinkedList[int]()

	obj.InsertFirst(51)
	obj.InsertFirst(21)
	obj.InsertFirst(11)
	obj.InsertLast(101)

	obj.Display()
	fmt.Printf("Number of elements are: %d\n", obj.Count())

    // Demonstration of Generics: Creating a list of strings
    strObj := NewSinglyLinkedList[string]()
    strObj.InsertFirst("World")
    strObj.InsertFirst("Hello")
    strObj.Display()
}
*/
