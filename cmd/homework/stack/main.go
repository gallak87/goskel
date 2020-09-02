package main

import (
	"errors"
	"fmt"
)

func main() {
	var s Stack = &stack{}

	fmt.Println("Push: 1, 2, 3")

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, _ := s.Pop()
	fmt.Println(fmt.Sprintf("Pop: %s", v))
	v, _ = s.Pop()
	fmt.Println(fmt.Sprintf("Pop: %s", v))
	v, _ = s.Pop()
	fmt.Println(fmt.Sprintf("Pop: %s", v))

}

type Stack interface {
	Push(v interface{})
	Pop() (interface{}, error)
	Size() int
	Clear()
}

type stack struct {
	sArr []interface{}
}

func (s *stack) Push(v interface{}) {
	s.sArr = append([]interface{}{v}, s.sArr...)
}

func (s *stack) Pop() (interface{}, error) {
	if s.Size() == 0 {
		return nil, errors.New("empty stack")
	}
	output := s.sArr[0]
	s.sArr = s.sArr[1:]
	return output, nil
}

func (s *stack) Size() int {
	return len(s.sArr)
}

func (s *stack) Clear() {
	s.sArr = []interface{}{}
}