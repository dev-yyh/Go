package stack

import (
	"log"
)

type ArrayStack struct {
	top int
	array []interface{}
}

//Init - Init stack
func (s *ArrayStack) Init (size int) bool {
	if cap(s.array) > 0 {
		log.Println("stack already init")
		return false
	}
	s.top = -1
	s.array = make([]interface{},size)
	return true
}

//IsEmpty - check if stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return s.top == -1
}

//IsFull - check if stack is full
func (s *ArrayStack) IsFull() bool {
	return s.top == cap(s.array)-1
}

//Peek - gives top element
func (s *ArrayStack) Peek() interface{} {
	if s.IsEmpty() {
		log.Println("stack is empty")
		return nil
	}
	return s.array[s.top]
}

//Push - pushes element into stack
func (s *ArrayStack) Push(data interface{}) {
	if s.IsFull() {
		log.Println("stack is full")
	} else {
		s.top ++
		s.array[s.top] = data
	}
}

//Pop - pops element from stack
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		log.Println("stack is empty")
		return nil
	} else {
		data := s.array[s.top]
		s.top --
		return data
	}
}