package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
// the key to add should not already be in the tree
// We compare the key with the reciever and go to the child on the left if smaller and to the right if larger.
func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)

		}
	}

}

// Search will take in a key value
// and RETURN true if there is a node with that value
// if a key value K exists in the tree return true or else return false
func (n *Node) Search(k int) bool {
	count++
	// if its not available then it returns false
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}

	// if its not larger or smaller but equal then return true
	return true
}

func main() {
	tree := &Node{Key: 100}

	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(203))
	fmt.Println(count)

}
