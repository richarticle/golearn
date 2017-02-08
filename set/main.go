package main

import (
	"fmt"
	"sync"
)

type Set struct {
	m map[int]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{m: map[int]bool{}}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Del(item int) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	} else {
		return false
	}
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func main() {
	s := NewSet()

	s.Add(1)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Del(1)

	fmt.Println("List:", s.List())
	fmt.Println("IsEmpty:", s.IsEmpty())
	fmt.Println("Has 1:", s.Has(1))

	s.Clear()
	fmt.Println("Clear set")
	fmt.Println("List:", s.List())
}
