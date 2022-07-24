package Set

import "fmt"

type void struct{}

var null = struct{}{}

type set struct {
	setmap map[int]void
}

func Set() *set {
	s := &set{}
	s.setmap = make(map[int]void)
	return s
}

func (s *set) Length() int {
	return len(s.setmap)
}

func (s *set) Add(val int) {
	s.setmap[val] = null
}

func (s *set) Contains(val int) bool {
	_, ok := s.setmap[val]
	return ok
}

func (s *set) Remove(val int) {
	delete(s.setmap, val)
}

func (s *set) Clear() {
	s.setmap = make(map[int]void)
}

func (s *set) Print() {
	fmt.Print("{")
	for k := range s.setmap {
		fmt.Print(k, ",")
	}
	fmt.Print("}\n")
}
