package main

import (
	"fmt"
	"time"
)

func main() {
	var err error

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLetters()
	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in doWork.")
}
