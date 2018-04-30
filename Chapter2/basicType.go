package main

import "fmt"

func main() {
	var i uint = 12345
	var i64 uint64 = uint64(i)
	fmt.Println(i64)

	var str string
	str = "あ"
	str += "いう"
	str += str + "エオ"
	fmt.Println(str)
}
