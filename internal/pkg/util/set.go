package util

type Set struct {
	set map[interface{}]struct{}
}

func (s *Set) IsEmpty() bool {
	return len(s.set) == 0
}

func (s *Set) Add(item interface{}) {
	s.set[item] = struct{}{}
}

func (s *Set) Remove(item interface{}) bool {
	if _, ok := s.set[item]; ok {
		delete(s.set, item)
		return true
	}
	return false
}

func (s *Set) Contains(item interface{}) bool {
	if _, ok := s.set[item]; ok {
		return true
	}
	return false
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) Clear() {
	clear(s.set)
}

func NewSet() *Set {
	return &Set{
		set: make(map[interface{}]struct{}),
	}
}
