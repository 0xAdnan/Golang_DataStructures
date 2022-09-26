package main

import "fmt"

// Stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push will add value at the end
// We take the stack as a pointer receiver because we would actually want to change the receiver
func (s *Stack) Push(i int) {
	// append this integer to two items, because append will return a slice with the appended integer we need to replace that with the existing items
	s.items = append(s.items, i)

}

// Pop will remove value at the end
// and RETURN the removed value
// In this we are not going to take in anything but we are going to return the integer that was removed
func (s *Stack) Pop() int {
	// [:len(s.items)-1] means its going to start from the beginnig and it's going to just leave one out at the end and replace items
	// we store the removed item in a variable first because we cannot return an item that has been removed already
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove

}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

}
