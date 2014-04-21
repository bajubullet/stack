package stack

import (
  "errors"
  "fmt"
)

type Stack struct {
  data []int
  top int
}

// Checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
  if s.top == 0 {
    return true
  }
  return false
}

// Checks if the stack is full.
func (s *Stack) IsFull() bool {
  if (s.top == len(s.data)-1) {
    return true
  }
  return false
}

// Pushes and item onto stack.
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

// Pops and an item from the stack.
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

// Print stack.
func (s *Stack) Print() {
  if s.IsEmpty() {
    fmt.Println("Empty stack.")
  } else {
    fmt.Println(s.data[:s.top])
  }
}

// Creates a new empty stack of given size.
func NewStack(size int) *Stack {
  s := new(Stack)
  s.data = make([]int, size)
  return s
}
