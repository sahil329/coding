package main

import "fmt"

type void struct{}

var null = struct{}{}

type Set struct {
	setmap map[int]void
}

func set() *Set {
	s := &Set{}
	s.setmap = make(map[int]void)
	return s
}

func (s *Set) length() int {
	return len(s.setmap)
}

func (s *Set) add(val int) {
	s.setmap[val] = null
}

func (s *Set) contains(val int) bool {
	_, ok := s.setmap[val]
	return ok
}

func (s *Set) remove(val int) {
	delete(s.setmap, val)
}

func (s *Set) clear() {
	s.setmap = make(map[int]void)
}

func (s *Set) print() {
	fmt.Print("{")
	for k := range s.setmap {
		fmt.Print(k, ",")
	}
	fmt.Print("}\n")
}

func main() {
	s := set()
	s.add(5)
	s.add(6)
	s.add(9)
	s.add(5)

	s.print()
	fmt.Print("Removing element 5 : ")
	s.remove(5)
	s.print()

	fmt.Println("Set contains 6? : ", s.contains(6))
	s.clear()
	fmt.Println("Clearning Set : ")
	s.print()

	s.add(1)
	s.add(3)
	s.add(2)
	s.add(3)
	fmt.Println("Length of set is : ", s.length())
	s.print()
}
