package main

import "fmt"

func main() {
	count := 50
	fibonacci := []int{1, 1}

	for i := 2; i < count; i++ {
		current := fibonacci[i-2] + fibonacci[i-1]
		fibonacci = append(fibonacci, current)
	}

	for i, v := range fibonacci {
		fmt.Printf("fib[%v]=%v\n", i, v)
	}
}
