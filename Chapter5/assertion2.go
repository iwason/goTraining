package main

import "fmt"

func main() {
	var i interface{} = "test"
	s1, ok := i.(string)
	fmt.Println(s1)

	s2, ok := i.(interface {
		dummy()
	})
	/**
		失敗する型変換の際に、第二引数を指定しないとランタイムパニックが発生する
		s3 := i.(interface {
			dummy()
		})
	**/
	fmt.Println(s2, ok)

	//	fmt.Println(s3)
}
