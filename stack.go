package stack

import "errors"

type Stack struct {
  data []int
  top int
}

func (s *Stack) IsEmpty() bool {
  if s.top == 0 {
    return true
  }
  return false
}

func (s *Stack) IsFull() bool {
  if (s.top == len(s.data)-1) {
    return true
  }
  return false
}

func (s *Stack) Push(num int) (err error) {
  err = nil
  if s.IsFull() {
    err = errors.New("Stack is aleady full.")
    return
  }
  s.data[s.top] = num
  s.top++
  return
}

func (s *Stack) Pop() (value int, err error) {
  err = nil
  value = 0
  if s.IsEmpty() {
    err = errors.New("Stack is aleady empty.")
  } else {
    s.top--
    value = s.data[s.top]
  }
  return
}

func NewStack(size int) *Stack {
  s := new(Stack)
  s.data = make([]int, size)
  return s
}
