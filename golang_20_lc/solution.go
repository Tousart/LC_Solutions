package main

import "fmt"

type Stack struct {
	values []rune
}

func (s *Stack) pop() rune {
	if len(s.values) == 0 {
		return -1
	}
	elem := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return elem
}

func isValid(s string) bool {
	var stack Stack
	for _, ph := range s {
		if ph == '(' || ph == '[' || ph == '{' {
			stack.values = append(stack.values, ph)
		} else if ph == ')' && stack.pop() != '(' {
			return false
		} else if ph == ']' && stack.pop() != '[' {
			return false
		} else if ph == '}' && stack.pop() != '{' {
			return false
		}
	}
	return len(stack.values) == 0
}

func main() {
	s := "}"
	fmt.Println(isValid(s))
}
