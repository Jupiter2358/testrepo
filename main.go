package main

import "fmt"

func main() {
	fmt.Printf("add(1,2) = %d\n", add(1,2))
}

func add(a, b int) int {
	return a + b
}