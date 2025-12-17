package main

import "errors"

type Stack struct {
	data     []string
	capacity int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		data:     make([]string, 0, capacity),
		capacity: capacity,
	}
}

func (s *Stack) Push(value string) error {
	if len(s.data) >= s.capacity {
		return errors.New("stack is full")
	}
	s.data = append(s.data, value)
	return nil
}

func (s *Stack) Pop() (string, error) {
	if len(s.data) == 0 {
		return "", errors.New("stack is empty")
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, nil
}

func (s *Stack) Peek() (string, error) {
	if len(s.data) == 0 {
		return "", errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) ToVector() []string {
	result := make([]string, len(s.data))
	copy(result, s.data)
	return result
}

func (s *Stack) FromVector(elements []string) {
	s.data = make([]string, 0, len(elements)+10)
	s.capacity = len(elements) + 10
	for _, element := range elements {
		s.data = append(s.data, element)
	}
}
