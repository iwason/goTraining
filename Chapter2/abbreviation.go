package main

import "fmt"

func abbreviation() {
	a, b := 1, 2
	a, b, c := 3, 4, 5
	d := ` string abbreviation
	with back quote.`
	fmt.Println(a, b, c, d)
}
