//Pointers stores the memory address of another variable.
//referencing $ and dereferencing *
//allow us to modify the value of a variable indirectly
//var ptr *int (declaring the pointer)
//there is no pointer arithmatic in go, unlike c


package main

import "fmt"

func main() {
	
	var ptr *int
	var a int = 10
	ptr = &a //referecing a pointer
		

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(&a)	  //same as ptr
	fmt.Println(*ptr) //dereference the pointer
	//a pointer that is not popointing to any value is a nil pointer (when no value is declared)
	if ptr == nil {
		fmt.Println("Pointer is nil")
	}


	modifyValue(ptr)
	fmt.Println(a)
}




func modifyValue(ptr *int) {	//when we pass a variable to a function, a copy of that variable is passed to the func.
	*ptr++			//but when we are using pointers, we are accessing the memory address of the var
				//and the value is modified. Thus no copy. :)
}
