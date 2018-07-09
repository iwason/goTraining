package main

import "fmt"

func group() {
	//同一の値が定義される
	const (
		a = 1
		b
		c
	)
	//iota演算子
	const (
		d = iota
		e
		f
		g
	)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f, g)
}
