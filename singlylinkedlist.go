package main

import "fmt"

type node struct {
	data int
	// next = address of the next node, so it needs to be a pointer
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	// inside we make a temporary place called second and put the current head over there.
	second := l.head
	// set the new node as head
	l.head = n
	// let the new head point to the old head which is second
	l.head.next = second
	// we need to increase the length because we added something
	l.length++

}

// to print out the data of every node in the list
// in this we use value receiver and not pointer reciver because in this we will just print out stuff and not make changes to the receiver
func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

//a method to delete a node with a given value
func (l *linkedList) deleteWithValue(value int) {
	// handling an empty list error
	if l.length == 0 {
		return
	}
	// handling the header as the value error
	if l.head.data == value {
		// the second node should become the header
		l.head = l.head.next
		l.length--
		return

	}
	// we put the head node into a variable to delete
	previousToDelete := l.head
	// we are not going to compare the data of the node to delete, we are going to compare the data of the next node

	for previousToDelete.next.data != value {
		// handling an error where the value is not available in the list
		if previousToDelete.next.next == nil {
			return
		}
		// to skip the node with the input value, we need to make modifications in the previous nodes next, so that it skips the node we want to delete
		// since this is a singly linked list we do not have info of the previous node address, thats why we are comparing the data in the next node.
		// so this is going to loop until it finds the match
		previousToDelete = previousToDelete.next
	}
	// updating the pointer so that it skips the node that needs to be deleted
	previousToDelete.next = previousToDelete.next.next
	l.length--

}

func main() {
	mylist := linkedList{}
	// we need to pass in a pointer thats what the & sign is there for
	node1 := &node{data: 58}
	node2 := &node{data: 48}
	node3 := &node{data: 98}
	node4 := &node{data: 42}
	node5 := &node{data: 18}
	// and then we are able to prepend it
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.printListData()
	mylist.deleteWithValue(98)
	mylist.printListData()

	// Some Special cases

	// Suppose we want to delete a value not in the linkedList, it will show runtime error
	mylist.deleteWithValue(100)

	// Suppose we want to delete the head value 18, it will show runtime error
	mylist.deleteWithValue(18)
	mylist.printListData()

	// Suppose we have an empty list and we try to delete a value, it will show runtime error
	emptyList := linkedList{}
	emptyList.deleteWithValue(98)

}
