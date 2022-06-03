package main

import (
	"fmt"
	"time"
)

type MultiPlyFunc func(int, int) int

func multiply(a, b int) int {
	return a * b
}

func execTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()
		c := f(a, b)
		end := time.Since(start)
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c
	}
}

func main() {
	value := execTime(multiply)(3, 11)
	fmt.Printf("this value is %d", value)
}
