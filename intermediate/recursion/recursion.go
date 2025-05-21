//Recursion is a function that calls itself.
//It breaks down a problem of the same type. Untill it becomes simple enough to solve directly.
//Every recursion has a base case where the recursion stops.

package main

import "fmt"

func main() {
	
	fmt.Println("factorial of n")
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
	

	fmt.Println("Sum of digits")
	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12345))


	fmt.Println("Fibonacci")
	
	for i:=10; i>0; i-- {
		fmt.Println(fibonacci(i))
	}
}


func factorial (n int) int {
	//Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	//Recusive case: factorial og n is n * factorial(n-1)
	return n * factorial(n-1)
	//n * (n-1) * factorial(n-2)
}


func sumOfDigits(n int) int {
	//Base case
	if n<10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
