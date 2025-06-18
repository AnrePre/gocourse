package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, "requestID", "sdaf43348958u")

	go doWork(ctx)
	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}
}

//func checkEvenOdd(ctx context.Context, num int) string {
//	select {
//	case <-ctx.Done():
//		return "Operation canceled"
//	default:
//		if num%2 == 0 {
//			return fmt.Sprintf("%d is even", num)
//		} else {
//			return fmt.Sprintf("%d is odd", num)
//		}
//	}
//}
//
//func main() {
//	ctx := context.TODO()
//
//	result := checkEvenOdd(ctx, 5)
//	fmt.Println("Result with context.TODO():", result)
//	ctx = context.Background()
//	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
//	defer cancel()
//
//	result = checkEvenOdd(ctx, 10)
//	fmt.Println("Result from timeout context:", result)
//
//	time.Sleep(2 * time.Second)
//	result = checkEvenOdd(ctx, 15)
//	fmt.Println("Result after timeout:", result)
//}
