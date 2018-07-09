package main

import "fmt"

func main() {
	fmt.Println("開始")
	defer delay()
	fmt.Println("終了")
}

func delay() {
	fmt.Println("delay呼び出し")
}
