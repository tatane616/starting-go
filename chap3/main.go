package main

import (
	"fmt"
	"runtime"
)

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func main() {
	q, r := div(19, 7)
	go fmt.Printf("商=%d 余剰=%d\n", q, r)
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
