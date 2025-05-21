//Create a Program that is run in the terminal. Where the user is taught the whole intermediate section of the go course. Thus add explanations and also examples of the code. This is a cool idea because when I get rusty on a section then I can just run the section that I would like to revise and then I will get the answers.

package main

import "fmt"

func main() {
	fmt.Println(".\n.\n.\nWelcome to My Golang Revision Tool.")
	fmt.Println("Last updated 05/09/2025")
	//Intro

	//Select Basics/Intermediate/Advanced
mainLoop:
	for {
		fmt.Println("Please Select 1.Basics or 2.Intermediate or 3.Advanced or 4.Exit")
		var mode int
		fmt.Scan(&mode)
		switch mode {
		case 1:
			fmt.Println("You have selected Basics")

		case 2:
			fmt.Println("You have selected Intermediate")
			intermediate()
		case 3:
			fmt.Println("You have selected Advanced")

		case 4:
			fmt.Println("You have selected Exit")
			break mainLoop
		default:
			fmt.Println("You have selected an invalid value, exiting...")
			break mainLoop
		}
	}
}

func intermediate() {
intermediateLoop:
	for {
		fmt.Println("\n---\nPlease select the chapter that you would like to revisit")
		fmt.Print("1. Closures\n2. Recursion\n3. Pointers\nYour answer: ")
		var answer int
		fmt.Scan(&answer)

		//Switch over choices
		switch answer {
		case 1:
			//closures
			fmt.Println("---\nClosures")
			fmt.Println("Definition: A closure in Go is a function that captures and uses variables from its surrounding scope, even after that scope has exited.")
		case 2:
			//recursion
			fmt.Println("---\nRecursion")
			fmt.Println("Recursion in Go is when a function calls itself to solve a problem by breaking it down into smaller sub-problems, usually with a base case to stop the calls.")
		case 3:
			//Pointers
			fmt.Println("---\nPointers")
			fmt.Println("Pointers in Go are variables that store the memory address of another variable, allowing functions to modify values directly.")
		case 4:
			fmt.Println("---\nStrings and Runes")

		case 5:
			fmt.Println()

		default:
			fmt.Println("Exit")
			break intermediateLoop
		}
		//Check if user wants to run the program agian or exit.
		var input string
		fmt.Print("Do you want to continue? (y/n): ")
		fmt.Scanln(&input)

		if input == "y" {
			fmt.Println("Continuing...")
		} else if input == "n" {
			fmt.Println("Exiting...")
			break // exit the loop
		} else {
			fmt.Println("Invalid input, please enter 'y' or 'n'")
		}
	}
}

func closureFunc() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Recursion function of n + (n-1)
func recursionFunc(n int) int {
	if n < 1 {
		return n
	}
	return n + recursionFunc(n-1)
}
