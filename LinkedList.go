package main

import "fmt"

type LNode struct {
	val  int
	next *LNode
}

func createLinkedList(l []int) *LNode {
	var head *LNode = &LNode{} // 1 : Initialized
	cur := head
	for _, n := range l {
		if head != nil {
			node := LNode{
				val:  n,
				next: &LNode{}, // 2 : Initialized
			}
			cur.next = &node
		} else {
			head.val = n
			cur = head
		}
		cur = cur.next
	}
	// make last node nil 'cause GOTO 2
	cur.next = nil

	// send from next node 'cause GOTO 1
	return head.next
}

func getLength(head *LNode) int {
	n := 0
	cur := head
	for cur != nil {
		cur = cur.next
		n++
	}
	return n
}

func (head *LNode) getLastNode() *LNode {
	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

func (head *LNode) insert(key int) {
	node := LNode{}
	node.val = key

	if head != nil {
		last := head.getLastNode()
		last.next = &node
	} else {
		head.next = &node
	}

}

func (head *LNode) deleteF() *LNode {
	head = head.next
	return head
}

func printLL(head *LNode) {
	cur := head
	for cur != nil {
		fmt.Print(cur.val)
		cur = cur.next
		if cur != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func main() {
	vals := []int{1, 2, 3, 4, 5}

	var root *LNode = createLinkedList(vals)
	printLL(root)
	root.insert(6)
	printLL(root)
	root = root.deleteF()
	printLL(root)

	r := &LNode{}
	for i := 1; i < 8; i++ {
		go r.insert(i)
	}

	printLL(r.next)
	// r.next = r.next.deleteF()
	// printLL(r.next)
}
