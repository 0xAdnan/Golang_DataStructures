package main

import "fmt"

//Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue adds a value at the end. Enqueue works same like stacks
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)

}

// Dequeue remove a value at the front
// and RETURNs the removed value
func (q *Queue) Dequeue() int {
	//  we will need to remove the first integer the one in the front index 0 and store it in a variable
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove

}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
