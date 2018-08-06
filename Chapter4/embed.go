package main

import "fmt"

type embedded struct {
	i int
}

func (value embedded) doSomething() {
	fmt.Println("do something")
}

type test struct {
	embedded
}

func main() {
	var x test
	x.doSomething()
	fmt.Println(x.i)
}
