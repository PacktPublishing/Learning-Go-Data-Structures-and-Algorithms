//+build disable

package main

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	left := len(s.items)

	if left == 0 {
		return -1
	}

	item, items := s.items[left-1], s.items[0:left-1]
	s.items = items
	return item
}

func main() {
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
	println(s.Pop())
}
