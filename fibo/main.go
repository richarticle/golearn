package main

import "fmt"

func main() {

	fmt.Println("Fibonacci sequence:")
	fmt.Println("Recursive: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fiboRecur(i))
	}
	fmt.Println("\nLoop: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fiboLoop(i))
	}
}

func fiboRecur(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fiboRecur(n-1) + fiboRecur(n-2)
}

func fiboLoop(n int) int {
	a, b := 0, 1
	for i := n; i > 0; i-- {
		a, b = b, a+b
	}
	return a
}
