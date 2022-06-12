package main

import (
	"fmt"
	"time"
)

type MultiPlyFunc func(int, int) int

func multiply(a, b int) int {
	return a * b
}

func execTime(function MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()
		call := function(a, b)
		end := time.Since(start)
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return call
	}
}

func main() {
	fmt.Printf("this value is %d", execTime(multiply)(3, 11))
}
