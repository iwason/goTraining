package main

import (
	"fmt"
)

func main() {
	x := "abcd"[1:4]
	//ひらがなはutf-8で3バイト
	y := "あいうえお"[3:9]
	//い が化ける
	z := "あいうえお"[0:4]

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
