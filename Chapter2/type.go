package main

import "fmt"

func integerConvert() {
	type myInteger int
	var i myInteger = 123
	fmt.Println(i)
}

type profile struct {
	height int
	weight int
}
