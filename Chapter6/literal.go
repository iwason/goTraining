package main

import "fmt"

/**
リテラルで初期化するときに配列のサイズを下回ると０で初期化される
*/
func main() {
	array1 := [6]float32{}
	array2 := [4]int{1, 2, 3}

	array3 := [...]string{
		"One",
		"Two",
		//初期化の際はカンマで終わらせない限り
		//閉じカッコを同じ行で書かないとコンパイルエラー
		"Three"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	fmt.Println(len(array3))
}
