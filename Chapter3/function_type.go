package main

import "fmt"

func main() {
	var f func(int, int) int

	f = func(a int, b int) int {
		return a + b
	}

	fmt.Println(f(3, 4))

	f = multiply
	fmt.Println(f(3, 4))
}

func multiply(a int, b int) int {
	return a * b
}
