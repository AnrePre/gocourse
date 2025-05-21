//write functions data structs and algorithms to handle types without defining the type.
//please note, that when a generic type is used for the first time, every other time where the same generic function or struct are used, it must be of the same type. Thus using it for the first time defines the type!!!
package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop()(T, bool) {
	if len(s.elements) == 0{
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}


func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}


func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("The stack is empty")
		return
	}
	fmt.Print("Stack elements: ")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println("")
}






func main() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)
	
	x1, y1 := "jane", "john"
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, y1)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	fmt.Println(intStack.isEmpty())
	intStack.push(3)
	fmt.Println(intStack.pop())
	intStack.printAll()
}



