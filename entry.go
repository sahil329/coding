package main

import (
	"fmt"

	"practice.com/Coding/LinkedList"
	"practice.com/Coding/Set"
)

func main() {
	vals := []int{1, 2, 3, 4, 5}

	var root *LinkedList.LNode = LinkedList.CreateLinkedList(vals)
	LinkedList.PrintLL(root)
	s := Set.Set()
	s.Add(5)
	s.Add(6)
	s.Add(9)
	s.Add(5)

	s.Print()
	fmt.Print("Removing element 5 : ")
	s.Remove(5)
	s.Print()

	fmt.Println("Set contains 6? : ", s.Contains(6))
	s.Clear()
	fmt.Println("Clearning Set : ")
	s.Print()

	s.Add(1)
	s.Add(3)
	s.Add(2)
	s.Add(3)
	fmt.Println("Length of set is : ", s.Length())
	s.Print()
}
