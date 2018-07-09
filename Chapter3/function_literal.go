package main

import "fmt"

func main() {
	val := 123
	var fn func(int, int) int

	func(i int) {
		fmt.Println(i * val)
	}(10)

	f := func(s string) {
		fmt.Println(s)
	}

	fn = func(a int, b int) int {
		return a + b
	}

	f("hoge")
	fmt.Println(fn(3, 5))
}
