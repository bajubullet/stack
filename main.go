package main

import (
  "fmt"
  "github.com/bajubullet/stack"
)


func main() {
  s := stack.NewStack(5)
  s.Print()
  s.Push(5)
  s.Push(10)
  s.Push(2)
  fmt.Println(s.Pop())
  s.Print()
}
