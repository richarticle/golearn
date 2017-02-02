package main

import "fmt"

func main() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 1; i <= 5; i++ {
		n, ok := s.Pop()
		if ok {
			fmt.Printf("%d ", n)
		}
	}
}

type Stack struct {
	store []int
}

func (s *Stack) Push(n int) {
	s.store = append(s.store, n)
}

func (s *Stack) Pop() (n int, ok bool) {
	if len(s.store) == 0 {
		ok = false
		return
	}

	n = s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]

	ok = true

	return
}
