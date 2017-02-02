package main

import (
	"fmt"
	"time"
)

func main() {
	target := 100000000
	t1 := time.Now()
	s := Sum(1, target)
	fmt.Printf("Sum(1, %d) = %d, Time: %v\n", target, s, time.Since(t1))

	t1 = time.Now()
	s = QuickSum(1, target)
	fmt.Printf("QuickSum(1, %d) = %d, Time: %v\n", target, s, time.Since(t1))
}

func Sum(start, end int) int {
	s := 0
	for i := start; i <= end; i++ {
		s += i
	}
	return s
}

func QuickSum(start, end int) int {
	batch := 100000
	sumCh := make(chan int)
	i := 0
	n := 0
	for i = start; i <= end-batch+1; i += batch {
		go sumRoutine(i, i+batch-1, sumCh)
		n++
	}
	s := 0
	for j := i; j <= end; j++ {
		s += j
	}
	for ; n > 0; n-- {
		s += <-sumCh
	}

	return s
}

func sumRoutine(start, end int, sumCh chan int) {
	s := 0
	for i := start; i <= end; i++ {
		s += i
	}
	sumCh <- s
}
