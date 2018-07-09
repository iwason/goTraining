package main

import "fmt"

func main() {
	var z myType = 123

	z.println()
}

type myType int

func (value myType) println() {
	fmt.Println(value)
}
